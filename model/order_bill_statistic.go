package model

import (
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
	"strconv"
)

func CreateOrderBillStatisticFactory(sqlType string) *OrderBillStatisticModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &OrderBillStatisticModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type OrderBillStatisticModel struct {
	*BaseModel
}

func (obsm *OrderBillStatisticModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_bill  ")

	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := obsm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, obsm.count()
}

func (obsm *OrderBillStatisticModel) GetAllBill(params map[string]interface{}) ([]map[string]interface{}, int) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select sn,start_time,end_time,sum(price) price,sum(commi_price)commi_price," +
		"sum(discount_price)discount_price,sum(bill_price)bill_price,sum(refund_price)refund_price,sum(refund_commi_price)refund_commi_price from es_bill")

	sn := params["sn"].(string)
	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)

	if sn != "" {
		sqlString.WriteString(fmt.Sprintf(" where sn = %s", sn))
	}
	sqlString.WriteString(" group by sn,start_time,end_time")

	if okPageNo && okPageSize {
		sqlString.WriteString(" limit ")
		sqlString.WriteString(strconv.Itoa(pageNo - 1))
		sqlString.WriteString(",")
		sqlString.WriteString(strconv.Itoa(pageSize))
	}

	rows := obsm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, len(tableData)
}

func (obsm *OrderBillStatisticModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_bill;"
	)

	err := obsm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
