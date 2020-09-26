package model

import (
	"log"
	"orange/utils/yml_config"
)

func CreateOrderStatisticFactory(sqlType string) *OrderStatisticModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &OrderStatisticModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("brandModel工厂初始化失败")
	return nil
}

type OrderStatisticModel struct {
	*BaseModel
}

func (osm *OrderStatisticModel) GetSalesMoneyTotal(start, end string) map[string]interface{} {
	sqlString := "SELECT SUM(o.`order_price`) AS receive_money,SUM(r.`refund_price`) AS refund_money " +
		"FROM es_sss_order_data o LEFT JOIN es_sss_refund_data r ON o.`sn` = r.`order_sn` WHERE o.`create_time` >= ? AND o.`create_time` <= ? "

	rows := osm.QuerySql(sqlString, start, end)
	defer rows.Close()

	if rows != nil {
		for rows.Next() {
			// TODO
		}
		_ = rows.Close()
	}
	return nil
}
