package model

import (
	"log"
	"orange/utils/yml_config"
)

func CreateGoodsFactory(sqlType string) *GoodsModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &GoodsModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("goodsModel工厂初始化失败")
	return nil
}

type GoodsModel struct {
	*BaseModel
	RoleId    string   `json:"role_id"`
	RoleName  string   `json:"role_name"`
	AuthIds   string   `json:"-"`
	ExAuthIds []string `json:"auth_ids"`
}

func (u *GoodsModel) NewGoods(length int) {
	//sql := "select * from es_goods where market_enable = 1 and disabled = 1 order by create_time desc limit 0, ?"
	//rows := u.QuerySql(sql, length)
	//if rows != nil {
	//	tmp := make([]UsersModel, 0)
	//	for rows.Next() {
	//		err := rows.Scan(&u.Id, &u.UserName, &u.Department)
	//		if err == nil {
	//			tmp = append(tmp, *u)
	//		} else {
	//			log.Println("sql查询错误", err.Error())
	//		}
	//	}
	//	_ = rows.Close()
	//	return tmp
	//}
}
