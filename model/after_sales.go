package model

import (
	"Goshop/global/consts"
	"Goshop/utils/sql_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"

	"github.com/gogf/gf/encoding/gjson"
)

func CreateAfterSalesFactory(sqlType string) *AfterSalesModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &AfterSalesModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type AfterSalesModel struct {
	*BaseModel
}

func (asm *AfterSalesModel) List(params map[string]interface{}) ([]map[string]interface{}, int64) {
	baseSql := "select * from es_as_order where 1=1"
	return asm.buildData(baseSql, params)
}

func (asm *AfterSalesModel) buildSql(baseSql string, params map[string]interface{}) (string, string) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString(baseSql)

	pageNo, okPageNo := params["page_no"].(int)
	disabled, _ := params["disabled"].(string)
	memberId, okMemberId := params["member_id"].(string)
	sellerId, okSellerId := params["seller_id"].(string)
	keyword, okKeyword := params["keyword"].(string)
	serviceSn, okServiceSn := params["service_sn"].(string)
	orderSn, okOrderSn := params["order_sn"].(string)
	goodsName, okGoodsName := params["goods_name"].(string)
	serviceType, okServiceType := params["service_type"].(string)
	serviceStatus, okServiceStatus := params["service_status"].(string)
	startTime, okStartTime := params["start_time"].(string)
	endTime, okEndTime := params["end_time"].(string)
	createChannel, okCreateChannel := params["create_channel"].(string)
	pageSize, okPageSize := params["page_size"].(int)

	if disabled == "" {
		disabled = "NORMAL"
	}
	if disabled != "" {
		sqlString.WriteString(fmt.Sprintf(" where disabled = '%s'", disabled))
	}
	if memberId != "" && okMemberId {
		sqlString.WriteString(fmt.Sprintf(" and member_id = %s", memberId))
	}
	if sellerId != "" && okSellerId {
		sqlString.WriteString(fmt.Sprintf(" and seller_id = %s", sellerId))
	}

	if keyword != "" && okKeyword {
		sqlString.WriteString(fmt.Sprintf(" and (sn like '%s' or order_sn like '%s' or goods_json like '%s')",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%"))
	}

	if serviceSn != "" && okServiceSn {
		sqlString.WriteString(fmt.Sprintf(" and sn like '%s'", "%"+serviceSn+"%"))
	}

	if orderSn != "" && okOrderSn {
		sqlString.WriteString(fmt.Sprintf(" and order_sn like '%s'", "%"+orderSn+"%"))
	}
	if goodsName != "" && okGoodsName {
		sqlString.WriteString(fmt.Sprintf(" and goods_json like '%s'", "%"+goodsName+"%"))
	}

	if serviceType != "" && okServiceType {
		sqlString.WriteString(fmt.Sprintf(" and service_type = %s", serviceType))
	}

	if serviceStatus != "" && okServiceStatus {
		sqlString.WriteString(fmt.Sprintf(" and service_status = %s", serviceStatus))
	}

	if startTime != "" && okStartTime {
		sqlString.WriteString(fmt.Sprintf(" and create_time >= %s", startTime))
	}

	if endTime != "" && okEndTime {
		sqlString.WriteString(fmt.Sprintf(" and create_time <= %s", endTime))
	}

	if createChannel != "" && okCreateChannel {
		sqlString.WriteString(fmt.Sprintf(" and create_channel = %s", createChannel))
	}
	sqlString.WriteString(" order by create_time desc")

	sqlCountString := sqlString.String()

	if okPageNo && okPageSize {
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	return sqlString.String(), sqlCountString
}

func (asm *AfterSalesModel) buildData(baseSql string, params map[string]interface{}) ([]map[string]interface{}, int64) {
	sqlString, countSqlString := asm.buildSql(baseSql, params)

	rows := asm.QuerySql(sqlString)
	defer rows.Close()
	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil, 0
	}
	var recordList []map[string]interface{}
	for _, item := range tableData {
		var goodsList []interface{}

		p, err := gjson.DecodeToJson(item["goods_json"])
		if err != nil {
			item["goods_list"] = goodsList
		} else {
			item["goods_list"] = p.ToArray()
		}
		//获取售后服务生成的新订单编号
		newOrderSn := item["new_order_sn"].(string)
		//获取售后服务类型
		serviceType := item["service_type"].(string)
		// 获取售后服务状态
		serviceStatus := item["service_status"].(string)

		item["service_type_text"] = consts.ServiceTypeMap[serviceType]
		item["service_status_text"] = consts.ServiceTypeMap[serviceStatus]
		item["allowable"] = asm.ServiceOperateAllowable(newOrderSn, serviceType, serviceStatus)
		recordList = append(recordList, item)
	}
	return tableData, sql_utils.Count(countSqlString, asm.dbDriverRead)
}

func (asm *AfterSalesModel) ServiceOperateAllowable(orderSn, serviceType, serviceStatus string) map[string]bool {
	allowAllowable := map[string]bool{
		"allow_audit":               false,
		"allow_ship":                false,
		"allow_put_in_store":        false,
		"allow_admin_refund":        false,
		"allow_show_storage_num":    false,
		"allow_show_return_addr":    false,
		"allow_show_ship_info":      false,
		"allow_seller_refund":       false,
		"allow_seller_create_order": false,
		"allow_seller_close":        false,
	}

	//是否允许商家审核：售后服务状态为新申请即可被商家审核
	allowAllowable["allow_audit"] = serviceStatus == consts.ServiceStatusApply

	//是否允许用户退还商品（填充物流信息）：售后服务类型为退货并且售后服务状态为审核通过 或者 售后服务类型为换货并且售后服务状态为审核通过并且售后服务已成功生成了新订单
	allowAllowable["allow_audit"] = (consts.ServiceTypeReturnGoods == serviceType && serviceStatus == consts.ServiceStatusPass) ||
		(consts.ServiceTypeChangeGoods == serviceType && serviceStatus == consts.ServiceStatusPass && orderSn != "")

	//是否允许商家入库：售后服务类型为取消订单并且审核通过 或者 售后服务类型为退货或者换货并且用户已经返还了商品
	allowAllowable["allow_put_in_store"] = (consts.ServiceTypeOrderCancel == serviceType && serviceStatus == consts.ServiceStatusPass) ||
		((consts.ServiceTypeReturnGoods == serviceType || consts.ServiceTypeChangeGoods == serviceType) && serviceStatus == consts.ServiceStatusFullCourier)

	//是否允许平台进行退款：售后服务类型为退货或者取消订单 并且 售后服务状态为待人工处理
	allowAllowable["allow_admin_refund"] = (consts.ServiceTypeReturnGoods == serviceType || consts.ServiceTypeOrderCancel == serviceType) &&
		consts.ServiceStatusWaitForManual == serviceStatus

	//是否允许展示商品的入库数量：售后服务类型不为补发商品 并且 售后服务状态是已入库或者退款中或者待人工处理或者退款失败或者已完成
	allowAllowable["allow_show_storage_num"] = !(consts.ServiceTypeSupplyAgainGoods == serviceType) && (consts.ServiceStatusStockIn == serviceStatus ||
		consts.ServiceStatusREFUNDING == serviceStatus || consts.ServiceStatusWaitForManual == serviceStatus ||
		consts.ServiceStatusREFUNDFAIL == serviceStatus || consts.ServiceStatusCOMPLETED == serviceStatus)

	//是否允许展示退货地址：售后服务类型为退货或换货 并且 售后服务状态不等于待审核和审核未通过
	allowAllowable["allow_show_return_addr"] = (consts.ServiceTypeReturnGoods == serviceType || consts.ServiceTypeChangeGoods == serviceType) &&
		consts.ServiceStatusApply != serviceStatus && consts.ServiceStatusRefuse != serviceStatus

	//是否允许展示用户填写的物流信息：售后服务类型为退货或换货 并且 售后服务状态不等于待审核、审核未通过、审核通过、已关闭和异常状态
	allowAllowable["allow_show_ship_info"] = (consts.ServiceTypeReturnGoods == serviceType || consts.ServiceTypeChangeGoods == serviceType) &&
		consts.ServiceStatusApply != serviceStatus && consts.ServiceStatusPass != serviceStatus && consts.ServiceStatusRefuse != serviceStatus &&
		consts.ServiceStatusCLOSED != serviceStatus && consts.ServicestatuserrorException != serviceStatus

	//是否允许商家退款：售后服务类型为退货或取消订单 并且 售后服务状态等于已入库
	allowAllowable["allow_seller_refund"] = (consts.ServiceTypeReturnGoods == serviceType || consts.ServiceTypeOrderCancel == serviceType) &&
		consts.ServiceStatusStockIn == serviceStatus

	//是否允许商家手动创建新订单：售后服务类型为换货或补发商品 并且 售后服务状态等于异常状态
	allowAllowable["allow_seller_create_order"] = (consts.ServiceTypeSupplyAgainGoods == serviceType || consts.ServiceTypeChangeGoods == serviceType) &&
		consts.ServicestatuserrorException == serviceStatus

	//是否允许商家关闭售后服务单：售后服务类型为换货或补发商品 并且 售后服务状态等于异常状态
	allowAllowable["allow_seller_close"] = (consts.ServiceTypeSupplyAgainGoods == serviceType || consts.ServiceTypeChangeGoods == serviceType) &&
		consts.ServicestatuserrorException == serviceStatus
	return allowAllowable
}
