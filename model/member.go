package model

import (
	"log"
	"orange/utils/sql_utils"
	"orange/utils/yml_config"
)

func CreateMemberFactory(sqlType string) *MemberModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &MemberModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("goodsModel工厂初始化失败")
	return nil
}

type BaseMember struct {
	Uname      string `json:"uname"`
	Email      string `json:"email"`
	CreateTime string `json:"create_time"`
	Mobile     string `json:"mobile"`
	Nickname   string `json:"nickname"`
}

type MemberModel struct {
	*BaseModel
}

func (mm *MemberModel) NewMember(length int) (allMemberList []BaseMember) {
	var (
		sqlString = "select uname,email,create_time,mobile,nickname " +
			"from es_member order by create_time desc limit 0,?"
	)

	rows := mm.QuerySql(sqlString, length)
	defer rows.Close()

	if rows != nil {
		for rows.Next() {
			member := BaseMember{}
			err := sql_utils.ParseToStruct(rows, &member)
			if err != nil {
				log.Println("sql_utils.ParseToStruct 错误.", err.Error())
			}
			allMemberList = append(allMemberList, member)
		}
		_ = rows.Close()
	}
	return allMemberList
}
