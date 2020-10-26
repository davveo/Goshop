package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
)

func CreateWithDrawFactory(sqlType string) *WithDrawModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &WithDrawModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type WithDrawModel struct {
	*BaseModel
}

func (wdm *WithDrawModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_withdraw_apply")

	pageNo, okPageNo := params["page_no"].(int)
	status, okStatus := params["status"].(string)
	pageSize, okPageSize := params["page_size"].(int)
	endTime, okEndTime := params["end_time"].(string)
	startTime, okstartTime := params["start_time"].(string)
	uname, okuname := params["uname"].(string)

	if uname != "" && okuname {
		sqlString.WriteString(fmt.Sprintf(" and member_name like  '%s'", "%"+uname+"%"))
	}

	if startTime != "" && okstartTime {
		sqlString.WriteString(fmt.Sprintf(" and apply_time > '%s'", startTime))
	}

	if endTime != "" && okEndTime {
		sqlString.WriteString(fmt.Sprintf(" and apply_time < %s", endTime))
	}

	if status != "" && okStatus {
		sqlString.WriteString(fmt.Sprintf(" and status = '%s'", status))
	}

	sqlString.WriteString(" order by id desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := wdm.QuerySql(sqlString.String())
	defer rows.Close()

	ItemList, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return ItemList, wdm.count()
}

func (wdm *WithDrawModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_withdraw_apply"
	)

	err := wdm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
