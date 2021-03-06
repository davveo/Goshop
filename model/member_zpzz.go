package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"log"
)

func CreateMemberZpzzFactory(sqlType string) *MemberZpzzModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &MemberZpzzModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type MemberZpzzModel struct {
	*BaseModel
}

func (mzm *MemberZpzzModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_member_zpzz where id != 0 ")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	sqlString.WriteString(" order by apply_time desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := mzm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, mzm.count()
}

func (mzm *MemberZpzzModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_member_zpzz;"
	)

	err := mzm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
