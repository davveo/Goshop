package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"log"
)

func CreateAfterSaleChangeFactory(sqlType string) *AfterSaleChangeModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &AfterSaleChangeModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("AfterSaleChangeModel工厂初始化失败")
	return nil
}

type AfterSaleChangeModel struct {
	*BaseModel
}

func (asc *AfterSaleChangeModel) getModel(serviceSn string) map[string]interface{} {
	rows := asc.QuerySql("select * from es_as_change where service_sn = ?", serviceSn)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil
	}
	var tmp map[string]interface{}
	if len(tableData) > 0 {
		tmp = tableData[0]
	}
	return tmp
}
