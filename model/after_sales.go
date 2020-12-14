package model

import (
	"Goshop/global/consts"
	"Goshop/utils/sql_utils"
	"Goshop/utils/time_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"errors"
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
			item["goodsList"] = goodsList
		} else {
			item["goodsList"] = p.ToArray()
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

func (asm *AfterSalesModel) Detail(serviceSn string) (map[string]interface{}, error) {
	if serviceSn == "" {
		return nil, errors.New("售后服务单编号不能为空")
	}
	//根据售后服务单号获取服务单信息
	afterSale := asm.getService(serviceSn)
	if afterSale == nil {
		return nil, errors.New("售后服务单信息不存在")
	}
	//获取申请售后的订单信息
	orderSn, _ := afterSale["order_sn"].(string)
	order := CreateOrderFactory("").getOrder(orderSn)
	if order == nil {
		return nil, errors.New("订单信息不存在")
	}

	//如果售后服务类型为退货或取消订单，则需要获取退款账户相关信息
	serviceType, _ := afterSale["service_type"].(string)
	if serviceType == consts.ServiceTypeReturnGoods || serviceType == consts.ServiceTypeOrderCancel {
		afterSaleRefund := CreateAfterSalesRefundFactory("").getAfterSaleRefund(serviceSn)
		afterSale["refund_info"] = afterSaleRefund
	}
	//获取售后服务单允许操作情况
	serviceStatus, _ := afterSale["service_status"].(string)
	allowable := asm.ServiceOperateAllowable(orderSn, serviceType, serviceStatus)
	afterSale["allowable"] = allowable

	//获取申请售后的商品信息集合
	goodsList := CreateAfterSaleGoodsFactory("").listGoods(serviceSn)
	afterSale["goods_list"] = goodsList

	//获取售后服务收货地址相关信息
	afterSaleChange := CreateAfterSaleChangeFactory("").getModel(serviceSn)
	afterSale["change_info"] = afterSaleChange

	//获取售后服务物流相关信息
	express := asm.getExpress(serviceSn)
	afterSale["express_info"] = express

	//获取售后服务用户上传的图片信息
	afterSaleImages := CreateAfterSaleGalleryFactory("").listImages(serviceSn)
	afterSale["images"] = afterSaleImages

	//获取售后服务日志相关信息
	afterSaleLogs := CreateAfterSaleLogFactory("").listLogs(serviceSn)
	afterSale["logs"] = afterSaleLogs

	//获取平台所有的正常开启使用的物流公司信息集合
	logiList := CreateLogisticsCompanyFactory("").listAllNormal()
	afterSale["logi_list"] = logiList

	shipStatus, _ := order["ship_status"].(string)
	paymentType, _ := order["payment_type"].(string)

	//获取订单的发货状态
	afterSale["order_ship_status"] = shipStatus
	//获取订单的付款类型
	afterSale["order_payment_type"] = paymentType

	//如果退货地址为空，那么需要获取商家店铺的默认地址作为退货地址
	returnAddr, okReturnAddr := afterSale["return_addr"].(string)
	sellerId, _ := afterSale["seller_id"].(string)
	if returnAddr == "" && okReturnAddr {
		shopDetail := CreateShopFactory(nil, "").getShopDetail(sellerId)
		shopAdd, _ := shopDetail["shop_add"].(string)
		shopCity, _ := shopDetail["shop_city"].(string)
		shopTown, _ := shopDetail["shop_town"].(string)
		linkName, _ := shopDetail["link_name"].(string)
		linkPhone, _ := shopDetail["link_phone"].(string)
		shopCounty, _ := shopDetail["shop_county"].(string)
		shopProvince, _ := shopDetail["shop_province"].(string)
		returnAddr = fmt.Sprintf("收货人: %s, 联系方式: %s, 地址: %s",
			linkName, linkPhone, shopProvince+shopCity+shopCounty+shopTown+"  "+shopAdd)
		afterSale["return_addr"] = returnAddr
	}
	return afterSale, nil
}

func (asm *AfterSalesModel) ExportAfterSale(params map[string]interface{}) []map[string]interface{} {
	var sqlString bytes.Buffer
	sqlString.WriteString("select ao.sn as service_sn,ao.order_sn,ao.create_time,ao.member_name,ao.seller_name,ao.mobile,ao.service_type,ao.service_status, " +
		"ao.reason,ao.problem_desc,ao.apply_vouchers,ao.audit_remark,ao.stock_remark,ao.refund_remark,ar.refund_price,ar.agree_price,ar.actual_price,ar.refund_time, " +
		"ar.refund_way,ar.account_type,ar.return_account,ar.bank_name,ar.bank_account_number,ar.bank_account_name,ar.bank_deposit_name from es_as_order ao left join " +
		"es_as_refund ar on ao.sn = ar.service_sn where 1=1 ")

	disabled, _ := params["disabled"].(string)
	endTime, okEndTime := params["end_time"].(string)
	orderSn, okOrderSn := params["order_sn"].(string)
	sellerId, okSellerId := params["seller_id"].(string)
	startTime, okStartTime := params["start_time"].(string)
	serviceSn, okServiceSn := params["service_sn"].(string)
	goodsName, okGoodsName := params["goods_name"].(string)
	serviceType, okServiceType := params["service_type"].(string)
	serviceStatus, okServiceStatus := params["service_status"].(string)

	if disabled == "" {
		disabled = "NORMAL"
	}
	if disabled != "" {
		sqlString.WriteString(fmt.Sprintf(" where ao.disabled = '%s'", disabled))
	}

	if sellerId != "" && okSellerId {
		sqlString.WriteString(fmt.Sprintf(" and ao.seller_id = %s", sellerId))
	}

	if serviceSn != "" && okServiceSn {
		sqlString.WriteString(fmt.Sprintf(" and ao.sn like '%s'", "%"+serviceSn+"%"))
	}

	if orderSn != "" && okOrderSn {
		sqlString.WriteString(fmt.Sprintf(" and ao.order_sn like '%s'", "%"+orderSn+"%"))
	}

	if goodsName != "" && okGoodsName {
		sqlString.WriteString(fmt.Sprintf(" and ao.goods_json like '%s'", "%"+goodsName+"%"))
	}

	if serviceType != "" && okServiceType {
		sqlString.WriteString(fmt.Sprintf(" and ao.service_type = %s", serviceType))
	}

	if serviceStatus != "" && okServiceStatus {
		sqlString.WriteString(fmt.Sprintf(" and ao.service_status = %s", serviceStatus))
	}

	if startTime != "" && okStartTime {
		sqlString.WriteString(fmt.Sprintf(" and ao.create_time >= %s", startTime))
	}

	if endTime != "" && okEndTime {
		sqlString.WriteString(fmt.Sprintf(" and ao.create_time <= %s", endTime))
	}

	sqlString.WriteString(" order by ao.create_time desc")

	rows := asm.QuerySql(sqlString.String())
	defer rows.Close()
	tableData, err := sql_utils.ParseJSON(rows)
	if err != nil {
		log.Println("sql_utils.ParseJSON 错误", err.Error())
		return nil
	}
	var recordList []map[string]interface{}
	for _, item := range tableData {
		serviceSn := item["service_sn"].(string)
		refundWay, okRefundWay := item["refund_way"].(string)             // 退款方式
		serviceType, okServiceType := item["service_type"].(string)       //获取售后服务类型
		accountType, okAccountType := item["account_type"].(string)       // 账户类型
		serviceStatus, okServiceStatus := item["service_status"].(string) // 获取售后服务状态
		if refundWay != "" && okRefundWay {
			item["refund_way_text"] = consts.RefundWayMap[refundWay]
		}
		if accountType != "" && okAccountType {
			item["account_type_text"] = consts.AccountTypeMap[accountType]
		}
		if serviceType != "" && okServiceType {
			item["service_type_text"] = consts.ServiceTypeMap[serviceType]
		}
		if serviceStatus != "" && okServiceStatus {
			item["service_status_text"] = consts.ServiceTypeMap[serviceStatus]
		}

		goodsList := CreateAfterSaleGoodsFactory("").listGoods(serviceSn)
		goodsInfo := ""
		for _, goods := range goodsList {
			goodsName := goods["goods_name"].(string)
			price := goods["price"].(string)
			shipNum := goods["ship_num"].(string)
			returnNum := goods["return_num"].(string)
			storageNum, okStorageNum := goods["storage_num"].(string)
			if storageNum == "" && okStorageNum {
				storageNum = "未入库"
			}
			goodsInfo += "【商品名称：" + goodsName + "，商品价格：" + price + "，购买数量：" + shipNum +
				"，申请售后数量：" + returnNum + "，入库数量：" + storageNum + "】"
		}
		item["goods_info"] = goodsInfo

		// 组合收货地址信息
		change := CreateAfterSaleChangeFactory("").getModel(serviceSn)
		province := change["province"].(string)
		city := change["city"].(string)
		county := change["county"].(string)
		town := change["town"].(string)
		shipAddr := change["ship_addr"].(string)
		shipName := change["ship_name"].(string)
		shipMobile := change["ship_mobile"].(string)
		rogInfo := "收货地址：" + province + city + county + town + shipAddr + "，收货人：" + shipName + "，联系方式：" + shipMobile
		item["rog_info"] = rogInfo

		//组合用户退还商品的物流信息
		express := asm.getExpress(serviceSn)
		if express != nil {
			var shipTimeStr string
			shipTime, okShipTime := express["ship_time"].(int64)
			courierNumber := express["courier_number"].(string)
			courierCompany := express["courier_company"].(string)
			if okShipTime && shipTime != 0 {
				shipTimeStr = time_utils.FormatTimestamp(shipTime)
			}
			expressInfo := "物流公司：" + courierCompany + "，快递单号：" + courierNumber + "，发货时间：" + shipTimeStr
			item["express_info"] = expressInfo
		}

		recordList = append(recordList, item)
	}
	return tableData
}

func (asm *AfterSalesModel) getAfterSaleCount(memberId, sellerId string) int64 {
	var sqlString bytes.Buffer
	sqlString.WriteString(fmt.Sprintf("select count(*) from es_as_order where service_status != %s and service_status != %s ",
		consts.ServiceStatusCOMPLETED, consts.ServiceStatusRefuse))

	if memberId != "" {
		sqlString.WriteString(fmt.Sprintf(" and member_id = %s", memberId))
	}
	if sellerId != "" {
		sqlString.WriteString(fmt.Sprintf(" and seller_id = %s", sellerId))
	}

	return sql_utils.Count(sqlString.String(), asm.dbDriverRead)
}

func (asm *AfterSalesModel) getService(serviceSn string) map[string]interface{} {
	rows := asm.QuerySql("select * from es_as_order where sn = ?", serviceSn)
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

func (asm *AfterSalesModel) getCancelService(orderSn string) map[string]interface{} {
	rows := asm.QuerySql(
		"select * from es_as_order where order_sn = ? and service_type = ? and service_status != ? and service_status != ?",
		orderSn, consts.ServiceTypeOrderCancel, consts.ServiceStatusRefuse, consts.ServiceStatusCLOSED)
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

func (asm *AfterSalesModel) getExpress(serviceSn string) map[string]interface{} {
	rows := asm.QuerySql("select * from es_as_express where service_sn = ?", serviceSn)
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

func (asm *AfterSalesModel) getOrderItems(orderSn string, skuId int64) map[string]interface{} {
	rows := asm.QuerySql("select * from es_order_items where order_sn = ? and product_id = ?", orderSn, skuId)
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
