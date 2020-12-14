package model

import (
	"Goshop/global/consts"
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/gogf/gf/encoding/gjson"
)

func CreateAfterSalesRefundFactory(sqlType string) *AfterSalesRefundModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &AfterSalesRefundModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type AfterSalesRefundModel struct {
	*BaseModel
}

func (asfm *AfterSalesRefundModel) List(params map[string]string) ([]map[string]interface{}, int64) {
	var sqlString bytes.Buffer

	sqlString.WriteString("select * from es_refund")
	pageNo, pageSize := params["page_no"], params["page_size"]
	if disabled, ok := params["disabled"]; ok && disabled != "" {
		sqlString.WriteString(fmt.Sprintf(" where disabled = '%s'", disabled))
	} else {
		sqlString.WriteString(fmt.Sprintf(" where disabled = '%s'", "NORMAL"))
	}
	if memberId, ok := params["member_id"]; ok && memberId != "" {
		sqlString.WriteString(fmt.Sprintf(" and member_id = '%s'", memberId))
	}

	if sellerId, ok := params["seller_id"]; ok && sellerId != "" {
		sqlString.WriteString(fmt.Sprintf(" and seller_id = '%s'", sellerId))
	}

	if keyword, ok := params["keyword"]; ok && keyword != "" {
		sqlString.WriteString(fmt.Sprintf(" and (sn like '%s' or order_sn like '%s' or goods_json like '%s')",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%"))
	}
	if serviceSn, ok := params["service_sn"]; ok && serviceSn != "" {
		sqlString.WriteString(fmt.Sprintf(" and sn like '%s'", "%"+serviceSn+"%"))
	}
	if goodsName, ok := params["goods_name"]; ok && goodsName != "" {

		sqlString.WriteString(fmt.Sprintf(" and goods_json like '%s'", "%"+goodsName+"%"))
	}

	if orderSn, ok := params["order_sn"]; ok && orderSn != "" {
		sqlString.WriteString(fmt.Sprintf(" and order_sn like '%s'", "%"+orderSn+"%"))
	}
	if refundStatus, ok := params["refund_status"]; ok && refundStatus != "" {
		sqlString.WriteString(fmt.Sprintf(" and refund_status = '%s'", refundStatus))
	}
	if refundWay, ok := params["refund_way"]; ok && refundWay != "" {
		sqlString.WriteString(fmt.Sprintf(" and refund_way = '%s'", refundWay))
	}
	if startTime, ok := params["start_time"]; ok && startTime != "" {
		sqlString.WriteString(fmt.Sprintf(" and create_time >= '%s'", startTime))
	}
	if endTime, ok := params["end_time"]; ok && endTime != "" {
		sqlString.WriteString(fmt.Sprintf(" and create_time <= '%s'", endTime))
	}
	if createChannel, ok := params["create_channel"]; ok && createChannel != "" {
		sqlString.WriteString(fmt.Sprintf(" and create_channel = '%s'", createChannel))
	}

	sqlString.WriteString(" order by create_time desc")

	if pageNo != "" && pageSize != "" {
		pageNo, _ := strconv.Atoi(pageNo)
		pageSize, _ := strconv.Atoi(pageSize)
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	resSqlString := sqlString.String()
	rows := asfm.QuerySql(resSqlString)
	defer rows.Close()

	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}

	for _, item := range tableData {
		var goodsList []interface{}
		refundStatus := item["refund_status"].(string)
		if agreePrice, okAgreePrice := item["agree_price"].(float64); okAgreePrice {
			item["agree_price"] = agreePrice
		} else {
			item["agree_price"] = 0
		}

		if actualPrice, okActualPrice := item["actual_price"].(float64); okActualPrice {
			item["actual_price"] = actualPrice
		} else {
			item["actual_price"] = 0
		}

		p, err := gjson.DecodeToJson(item["goods_json"])
		if err != nil {
			item["goodsList"] = goodsList
		} else {
			item["goodsList"] = p.ToArray()
		}

		item["service_sn"] = item["sn"]
		item["refund_status_text"] = consts.RefundStatusMap[refundStatus]
	}

	return tableData, asfm.count(resSqlString)
}

func (asfm *AfterSalesRefundModel) count(SqlString string) (rows int64) {
	err := asfm.QueryRow(SqlString).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}

func (asfm *AfterSalesRefundModel) getAfterSaleRefund(serviceSn string) map[string]interface{} {
	rows := asfm.QuerySql("select * from es_as_refund where service_sn = ?", serviceSn)
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
