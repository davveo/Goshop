package model

import (
	"Goshop/utils/sql_utils"
	su "Goshop/utils/syncopate_utils"
	"Goshop/utils/yml_config"
	"log"
)

func CreateOrderStatisticFactory(sqlType string) *OrderStatisticModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &OrderStatisticModel{
			BaseModel: dbDriver,
			SyncUtil:  su.SyncopateUtil{},
		}
	}
	log.Fatal("brandModel工厂初始化失败")
	return nil
}

type OrderStatisticModel struct {
	*BaseModel
	SyncUtil su.SyncopateUtil
}

func (osm *OrderStatisticModel) GetSalesMoneyTotal(year, start, end string) map[string]interface{} {
	sqlString := "SELECT SUM(o.`order_price`) AS receive_money,SUM(r.`refund_price`) AS refund_money " +
		"FROM es_sss_order_data o LEFT JOIN es_sss_refund_data r ON o.`sn` = r.`order_sn` WHERE o.`create_time` >= ? AND o.`create_time` <= ? "

	sqlString = osm.SyncUtil.HandleSql(year, sqlString)
	rows := osm.QuerySql(sqlString, start, end)
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
