package model

import (
	"bytes"
	"log"
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"strconv"
)

func CreateTradePayLogFactory(sqlType string) *TradePayLogModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &TradePayLogModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type TradePayLogModel struct {
	*BaseModel
}

func (tplm *TradePayLogModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)
	sqlString.WriteString("select * from es_pay_log ")

	payWay := params["pay_way"].(string)
	endTime := params["end_time"].(string)
	orderSn := params["order_sn"].(string)
	payType := params["pay_type"].(string)
	startTime := params["start_time"].(string)
	payStatus := params["pay_status"].(string)
	pageNo, okPageNo := params["page_no"].(int)
	pageSize, okPageSize := params["page_size"].(int)
	payMemberName := params["pay_member_name"].(string)

	if payWay != "" {
		sqlString.WriteString(" and pay_way = ")
		sqlString.WriteString(payWay)
	}

	if payStatus != "" {
		sqlString.WriteString(" and pay_status = ? ")
		sqlString.WriteString(payStatus)
	}
	if startTime != "" && endTime != "" {
		sqlString.WriteString(" and pay_time BETWEEN ")
		sqlString.WriteString(startTime)
		sqlString.WriteString(" and ")
		sqlString.WriteString(endTime)
	}
	if payMemberName != "" {
		sqlString.WriteString(" and pay_member_name = ")
		sqlString.WriteString(payMemberName)
	}
	if orderSn != "" {
		sqlString.WriteString(" and order_sn like ")
		sqlString.WriteString(" '%" + orderSn + "%' ")
	}
	if payType != "" {
		var value string
		sqlString.WriteString(" and order_sn like ")
		if payType == "alipay" {
			value = "%支付宝%"
		} else if payType == "wechat" {
			value = "%微信%"
		}
		sqlString.WriteString(value)
	}

	sqlString.WriteString(" order by pay_log_id desc")

	if okPageNo && okPageSize {
		sqlString.WriteString(" limit ")
		sqlString.WriteString(strconv.Itoa(pageNo - 1))
		sqlString.WriteString(",")
		sqlString.WriteString(strconv.Itoa(pageSize))
	}

	rows := tplm.QuerySql(sqlString.String())
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	return tableData, tplm.count()
}

func (tplm *TradePayLogModel) count() (rows int64) {
	var (
		sql = "select count(*) from es_pay_log;"
	)

	err := tplm.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}
