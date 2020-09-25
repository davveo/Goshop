package model

import (
	"log"
	"orange/utils/yml_config"
)


func CreateCategoryFactory(sqlType string) *CategoryModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &CategoryModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("settingModel工厂初始化失败")
	return nil
}

type CategoryModel struct {
	*BaseModel
}