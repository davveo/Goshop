package model

import (
	"errors"
	"fmt"
	"log"
	"orange/global/variable"
	ojwt "orange/utils/jwt"
	"orange/utils/md5_encrypt"
	"orange/utils/yml_config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateUserFactory(sqlType string) *UsersModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &UsersModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("usersModel工厂初始化失败")
	return nil
}

type UsersModel struct {
	*BaseModel
	Id           int64  `json:"uid"`
	UserName     string `json:"username"`
	PassWord     string `json:"-"`
	RealName     string `json:"-"`
	Department   string `json:"department"`
	Face         string `json:"face"`
	Founder      int    `json:"founder"`
	RoleId       int    `json:"role_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserState    string `json:"user_state"`
}

func (u *UsersModel) Login(uuid, name, pass string) *UsersModel {
	sql := "select id, username, password, real_name, department, face, founder, role_id, user_state  " +
		"from es_admin_user where username = ? and password = ? and user_state = 0 limit 1"

	encyPassword := md5_encrypt.Base64Md5(pass)
	err := u.QueryRow(sql, name, encyPassword).Scan(
		&u.Id, &u.UserName, &u.PassWord, &u.RealName,
		&u.Department, &u.Face, &u.Founder, &u.RoleId, &u.UserState)

	if err == nil {
		// 账号密码验证成功
		if len(u.PassWord) > 0 && (u.PassWord == encyPassword) {
			u.AccessToken = u.createToken(variable.AccessTokenExpireTime)
			u.RefreshToken = u.createToken(variable.RefreshTokenExpireTime)

			accessTokenCacheKey := u.cacheName(variable.AccessTokenPrefix, string(u.Id), uuid)
			rds.Put(accessTokenCacheKey, u.AccessToken, variable.AccessTokenCacheExpireTime)

			refreshTokenCacheKey := u.cacheName(variable.RefreshTokenPrefix, string(u.Id), uuid)
			rds.Put(refreshTokenCacheKey, u.AccessToken, variable.RefreshTokenCacheExpireTime)

			adminDisabledCacheKey := u.cacheName(variable.AdminDisabledPrefix, string(u.Id))
			rds.Put(adminDisabledCacheKey, u.AccessToken, variable.AccessTokenCacheExpireTime)

			return u
		}
	} else {
		log.Println("根据账号查询单条记录出错:", err)
	}
	return nil
}

func (u *UsersModel) Logout(uuid, uid string) {

	accessTokenCacheKey := u.cacheName(variable.AccessTokenPrefix, uid, uuid)
	rds.Remove(accessTokenCacheKey)

	refreshTokenCacheKey := u.cacheName(variable.RefreshTokenPrefix, uid, uuid)
	rds.Remove(refreshTokenCacheKey)

	adminDisabledCacheKey := u.cacheName(variable.AdminDisabledPrefix, uid)
	rds.Remove(adminDisabledCacheKey)
}

func (u *UsersModel) Add(name string, pass string) *UsersModel {

	return nil
}

func (u *UsersModel) Edit(Id int64) (*UsersModel, error) {

	return nil, nil
}

func (u *UsersModel) Delete(Id int64) (bool, error) {
	sql := "update es_admin_user set user_state = -1 where id = ?"

	user := u.selectById(Id)
	if user == nil {
		return false, errors.New("当前管理员不存在")
	}
	// 校验要删除的管理员是否是最后一个超级管理员
	userList := u.list()
	if len(userList) <= 1 && u.Founder == 1 {
		return false, errors.New("必须保留一个超级管理员")
	}
	if u.ExecuteSql(sql, Id) > 0 {
		cacheKey := u.cacheName(variable.AdminDisabledPrefix, string(u.Id))
		rds.Put(cacheKey, -1, variable.AccessTokenCacheExpireTime)

		return true, nil
	}
	return false, errors.New("更新失败")
}

func (u *UsersModel) selectById(Id int64) *UsersModel {
	sql := "select * from es_admin_user where user_state = 0 and id = ?"

	rows := u.QuerySql(sql, Id)
	if rows != nil {
		for rows.Next() {
			err := rows.Scan(&u.Id, &u.UserName, &u.Department)
			if err == nil {
				return u
			}
		}
		_ = rows.Close()
	}
	return nil
}

func (u *UsersModel) createToken(expire int64) string {
	j := ojwt.NewJwt()
	claims := ojwt.CustomClaims{
		ID:       u.Id,
		UserName: u.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + 1000*expire,
			Issuer:    "orange",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return ""
	}
	return token
}

func (u *UsersModel) list() []UsersModel {
	sql := "select * from es_admin_user where founder = 1 and user_state = 0"

	rows := u.QuerySql(sql)
	if rows != nil {
		tmp := make([]UsersModel, 0)
		for rows.Next() {
			err := rows.Scan(&u.Id, &u.UserName, &u.Department)
			if err == nil {
				tmp = append(tmp, *u)
			} else {
				log.Println("sql查询错误", err.Error())
			}
		}
		_ = rows.Close()
		return tmp
	}
	return nil
}

func (u *UsersModel) exchangeToken(refreshToken string) string {
	return ""
}

func (u *UsersModel) cacheName(prefix string, params ...interface{}) string {
	return fmt.Sprintf(prefix, params...)
}
