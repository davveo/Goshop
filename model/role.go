package model

import (
	"Goshop/utils/yml_config"
	"encoding/json"
	"errors"
	"log"
)

func CreateRoleFactory(sqlType string) *RoleModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &RoleModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type RoleModel struct {
	*BaseModel
	RoleId    string   `json:"role_id"`
	RoleName  string   `json:"role_name"`
	AuthIds   string   `json:"-"`
	ExAuthIds []string `json:"auth_ids"`
}

type RoleTree struct {
	Title       string     `json:"title"`
	Identifier  string     `json:"identifier"`
	Checked     bool       `json:"checked"`
	AuthRegular string     `json:"authRegular"`
	Children    []RoleTree `json:"children"`
}

func (u *RoleModel) GetRoleMenu(Id int) ([]string, error) {
	var data []RoleTree

	role := u.GetRoleModel(Id)
	if role == nil {
		return nil, errors.New("此角色不存在")
	}

	err := json.Unmarshal([]byte(role.AuthIds), &data)
	if err != nil {
		return nil, errors.New("数据解析失败")
	}
	// 解析数据
	u.traveseTree(data)

	return u.ExAuthIds, nil
}

func (u *RoleModel) traveseTree(roleTree []RoleTree) {
	for _, v := range roleTree {
		if v.Checked {
			u.ExAuthIds = append(u.ExAuthIds, v.Identifier)
		}
		if len(v.Children) != 0 {
			u.traveseTree(v.Children)
		}
	}
}

func (u *RoleModel) GetRoleModel(Id int) *RoleModel {
	sql := "select role_id, role_name, auth_ids from es_role where role_id = ?"
	err := u.QueryRow(sql, Id).Scan(&u.RoleId, &u.RoleName, &u.AuthIds)
	if err == nil {
		return u
	}
	return nil
}
