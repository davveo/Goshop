package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
)

func CreateAfterSalesFactory(sqlType string) *AfterSalesModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &AfterSalesModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type AfterSalesModel struct {
	*BaseModel
}

func (asm *AfterSalesModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	baseSql := "select * from es_as_order "
	return asm.buildData(baseSql, params)
}

func (asm *AfterSalesModel) buildSql(baseSql string, params map[string]interface{}) (string, string) {
	var (
		sqlString bytes.Buffer
	)
	sqlString.WriteString(baseSql)
	pageNo, _ := params["page_no"].(int)
	disabled, _ := params["disabled"].(string)
	pageSize, _ := params["page_size"].(int)

	build := sql_utils.Builder{BaseSql: baseSql}

	if disabled == "" {
		disabled = "NORMAL"
	}
	if disabled != "" {
		sqlString.WriteString(fmt.Sprintf(" where disabled = '%s'", disabled))
	}

	sql := build.Where("disabled", disabled, "=").
		OrderBy("create_time", "desc")

	return sql.LimitOffset(pageNo, pageSize).ToString(), sql_utils.SqlCountString(sql.ToString())
}

func (asm *AfterSalesModel) buildData(baseSql string, params map[string]interface{}) ([]map[string]interface{}, int64) {
	sqlString, countSqlString := asm.buildSql(baseSql, params)

	rows := asm.QuerySql(sqlString)
	defer rows.Close()
	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}
	return tableData, sql_utils.Count(countSqlString, asm.dbDriverRead)
}
