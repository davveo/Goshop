package model

import (
	"Goshop/global/consts"
	"Goshop/model/com"
	"Goshop/utils/rabbitmq"
	"Goshop/utils/sql_utils"
	"Goshop/utils/time_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"errors"
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
	mq := rabbitmq.GetRabbitmq()
	if mq == nil {
		log.Fatal("goodsModel mq初始化失败")
	}
	amqpTemplate, err := mq.Producer("after-sales-refund")
	if err != nil {
		log.Fatal("goodsModel producer初始化失败")
	}

	if dbDriver != nil {
		return &AfterSalesRefundModel{
			BaseModel:    dbDriver,
			amqpTemplate: amqpTemplate,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type AfterSalesRefundModel struct {
	*BaseModel
	amqpTemplate *rabbitmq.Producer
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

func (asfm *AfterSalesRefundModel) AdminRefund(refundPrice float64, serviceSn, remark, serviceOperate string) error {
	if err := asfm.checkAdminRefund(refundPrice, serviceSn, remark); err != nil {
		return err
	}
	//获取售后服务单详细信息
	applyAfterSale, err := CreateAfterSalesFactory("").Detail(serviceSn)
	if err != nil {
		return err
	}
	//操作权限验证
	serviceType := applyAfterSale["service_type"].(string)
	serviceStatus := applyAfterSale["service_status"].(string)
	if !com.CheckOperate(serviceType, serviceStatus, serviceOperate) {
		return errors.New("当前售后服务单状态不允许进行退款操作")
	}

	refund := asfm.getModel(serviceSn)
	if refund == nil {
		return errors.New("售后退款单信息不存在")
	}
	// 获取退款时间
	refundTime := time_utils.CurrentTimeStamp()
	if asfm.ExecuteSql("update es_refund set refund_status = ?,refund_time = ?,actual_price = ? where sn = ?",
		consts.RefundStatusCompleted, refundTime, refundPrice, serviceSn) == -1 {
		return errors.New("更新退款失败")
	}

	//修改售后服务退款相关信息
	if asfm.ExecuteSql("update es_as_refund set actual_price = ?,refund_time = ? where service_sn = ?",
		refundPrice, refundTime, serviceSn) == -1 {
		return errors.New("更新退款售后失败")
	}

	//将售后服务单状态和退款备注
	if err = CreateAfterSaleOrderFactory("").updateServiceStatus(serviceSn, consts.ServiceStatusCOMPLETED,
		"", "", remark, ""); err != nil {
		return err
	}
	//发送售后服务完成消息
	afterSaleMsg := rabbitmq.BuildMsg(map[string]interface{}{
		"service_sn":     serviceSn,
		"service_type":   serviceType,
		"service_status": consts.ServiceStatusCOMPLETED,
	})
	err = asfm.amqpTemplate.Publish(consts.ExchangeAsStatusChange,
		consts.ExchangeAsStatusChange+"_QUEUE", afterSaleMsg)
	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	//新增退款操作日志
	logStr := "已成功将退款退还给买家，当前售后服务已完成。"
	if err = CreateAfterSaleLogFactory("").add(serviceSn, logStr, "系统"); err != nil {
		return err
	}
	return nil
}

func (asfm *AfterSalesRefundModel) checkAdminRefund(refundPrice float64, serviceSn, remark string) error {
	if serviceSn == "" {
		return errors.New("退款单编号不能为空")
	}
	if refundPrice <= 0 {
		return errors.New("退款金额不能小于或等于0元")
	}
	if remark != "" && len(remark) > 150 {
		return errors.New("退款备注不能超过150个字符")
	}
	return nil
}

func (asfm *AfterSalesRefundModel) getModel(serviceSn string) map[string]interface{} {
	rows := asfm.QuerySql("select * from es_refund where sn = ?", serviceSn)
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
