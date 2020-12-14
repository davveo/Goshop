package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"log"
)

func CreateAfterSaleGalleryFactory(sqlType string) *AfterSaleGalleryModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &AfterSaleGalleryModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("AfterSaleGalleryModel工厂初始化失败")
	return nil
}

type AfterSaleGalleryModel struct {
	*BaseModel
}

func (asgm *AfterSaleGalleryModel) listImages(serviceSn string) []map[string]interface{} {
	rows := asgm.QuerySql("select * from es_as_gallery where service_sn = ?", serviceSn)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil
	}

	return tableData
}
