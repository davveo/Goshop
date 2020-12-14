package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"log"
)

func CreateAfterSaleLogFactory(sqlType string) *AfterSaleLogModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &AfterSaleLogModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("AfterSaleLogModel工厂初始化失败")
	return nil
}

type AfterSaleLogModel struct {
	*BaseModel
}

func (asm *AfterSaleLogModel) listLogs(serviceSn string) []map[string]interface{} {
	rows := asm.QuerySql("select * from es_as_log where sn = ? order by log_time desc", serviceSn)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil
	}

	return tableData
}

func (asm *AfterSaleLogModel) add(serviceSn, logDetail, operator string) error {
	return nil
}
