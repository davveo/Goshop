package model

import (
	"log"
	"Goshop/utils/yml_config"
)

func CreateHealthFactory(sqlType string) *HealthModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &HealthModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type HealthModel struct {
	*BaseModel
}

func (h *HealthModel) Check() bool {
	return h.dbDriverRead.Ping() == nil
}
