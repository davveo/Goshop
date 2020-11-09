package model

import (
	"Goshop/utils/yml_config"
	"log"
)

func CreateHealthFactory(sqlType string) *HealthModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
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
	return h.dbDriverWrite.Ping() == nil
}
