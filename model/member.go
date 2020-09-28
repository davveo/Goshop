package model

import (
	"bytes"
	"log"
	"Eshop/utils/sql_utils"
	"Eshop/utils/yml_config"
	"strconv"
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

func (mm *MemberModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_member ")

	disabled := params["disabled"].(string)
	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	if disabled != "" {
		if disabled != "-1" && disabled != "0" {
			sqlString.WriteString(" where disabled = 0")
		} else {
			sqlString.WriteString(" where disabled = ")
			sqlString.WriteString(disabled)
		}
	} else {
		sqlString.WriteString(" where disabled = 0")
	}

	sqlString.WriteString(" order by create_time desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(" limit ")
		sqlString.WriteString(strconv.Itoa(pageNo - 1))
		sqlString.WriteString(",")
		sqlString.WriteString(strconv.Itoa(pageSize))
	}

	rows := mm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, mm.count()
}

func (mm *MemberModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_member"
	)

	err := mm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
