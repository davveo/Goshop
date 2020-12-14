package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"errors"
	"log"
)

func CreateAfterSaleGoodsFactory(sqlType string) *AfterSaleGoodsModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &AfterSaleGoodsModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("AfterSaleGoodsModel工厂初始化失败")
	return nil
}

type AfterSaleGoodsModel struct {
	*BaseModel
}

func (asgm *AfterSaleGoodsModel) listGoods(serviceSn string) []map[string]interface{} {
	rows := asgm.QuerySql("select * from es_as_goods where service_sn = ?", serviceSn)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil
	}

	return tableData
}

func (asgm *AfterSaleGoodsModel) updateStorageNum(serviceSn, skuId, num string) error {
	if asgm.ExecuteSql("update es_as_goods set storage_num = ? where service_sn = ? and sku_id = ?", num, serviceSn, skuId) == -1 {
		return errors.New("更新失败")
	}
	return nil
}
