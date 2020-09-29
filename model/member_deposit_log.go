package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"log"
)

func CreateMemberDepositLogFactory(sqlType string) *MemberDepositLogModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &MemberDepositLogModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type MemberDepositLogModel struct {
	*BaseModel
}

func (mdlm *MemberDepositLogModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_member_deposit_log ")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	sqlString.WriteString(" order by create_time desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := mdlm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, mdlm.count()
}

func (mdlm *MemberDepositLogModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_member_deposit_log;"
	)

	err := mdlm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
