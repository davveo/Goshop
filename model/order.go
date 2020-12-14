package model

import (
	"Goshop/global/consts"
	"Goshop/utils/sql_utils"
	"Goshop/utils/time_utils"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"log"
	"strconv"
)

func CreateOrderFactory(sqlType string) *OrderModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &OrderModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("healthModel工厂初始化失败")
	return nil
}

type OrderModel struct {
	*BaseModel
}

func (om *OrderModel) List(params map[string]string) ([]map[string]interface{}, int64) {
	var (
		sqlString bytes.Buffer
	)

	sqlString.WriteString("select * from es_order o where disabled=0 ")

	pageNo, pageSize := params["page_no"], params["page_size"]
	if keywords, ok := params["keyword"]; ok && keywords != "" {
		sqlString.WriteString(fmt.Sprintf(
			" and (sn like %s' or items_json like '%s' ) ",
			"%"+keywords+"%", "%"+keywords+"%"))
	}
	if sellerId, ok := params["seller_id"]; ok && sellerId != "" {
		sqlString.WriteString(fmt.Sprintf("  and o.seller_id = %s", sellerId))
	}
	if memberId, ok := params["member_id"]; ok && memberId != "" {
		sqlString.WriteString(fmt.Sprintf("  and o.member_id = %s", memberId))
	}
	if orderSn, ok := params["order_sn"]; ok && orderSn != "" {
		sqlString.WriteString(fmt.Sprintf(" and o.sn like '%s'", "%"+orderSn+"%"))
	}
	if tradeSn, ok := params["trade_sn"]; ok && tradeSn != "" {
		sqlString.WriteString(fmt.Sprintf(" and o.trade_sn like '%s'", "%"+tradeSn+"%"))
	}
	if startTime, ok := params["start_time"]; ok && startTime != "" {
		startTime, _ := strconv.ParseInt(startTime, 10, 64)
		sqlString.WriteString(fmt.Sprintf("  and o.create_time >= %s", time_utils.GetDayOfStart(startTime)))
	}
	if endTime, ok := params["end_time"]; ok && endTime != "" {
		endTime, _ := strconv.ParseInt(endTime, 10, 64)
		sqlString.WriteString(fmt.Sprintf("  and o.create_time <= %s", time_utils.GetDayOfEnd(endTime)))
	}
	if memberName, ok := params["member_name"]; ok && memberName != "" {
		sqlString.WriteString(fmt.Sprintf(" and o.member_name like '%s'", "%"+memberName+"%"))
	}
	if orderStatus, ok := params["order_status"]; ok && orderStatus != "" {
		sqlString.WriteString(fmt.Sprintf("  and o.order_status = %s", orderStatus))
	}
	if buyerName, ok := params["buyer_name"]; ok && buyerName != "" {
		sqlString.WriteString(fmt.Sprintf(" and o.items_json like '%s'", "%"+buyerName+"%"))
	}
	if goodsName, ok := params["goods_name"]; ok && goodsName != "" {
		sqlString.WriteString(fmt.Sprintf(" and o.items_json like '%s'", "%"+goodsName+"%"))
	}
	if shipName, ok := params["ship_name"]; ok && shipName != "" {
		sqlString.WriteString(fmt.Sprintf(" and o.ship_name like '%s'", "%"+shipName+"%"))
	}
	if paymentType, ok := params["payment_type"]; ok && paymentType != "" {
		sqlString.WriteString(fmt.Sprintf("  and o.payment_type = %s", paymentType))
	}
	if clientType, ok := params["client_type"]; ok && clientType != "" {
		sqlString.WriteString(fmt.Sprintf("  and o.client_type = %s", clientType))
	}
	if tag, ok := params["tag"]; ok && tag != "" {
		switch tag {
		case consts.OrderTagAll:
			break
		case consts.OrderTagWaitPay: // 待付款
			// 非货到付款的，未付款状态的可以结算 OR 货到付款的要发货或收货后才能结算
			str := " and ( ( ( payment_type!='cod' and  order_status='" + consts.OrderStatusConfirm + "') " +
				" or ( payment_type='cod' and   order_status='" + consts.OrderStatusRog + "'  ) ) " +
				" or order_status = '" + consts.OrderStatusNew + "' )"
			sqlString.WriteString(str)
			break
		case consts.OrderTagWaitShip:
			// 普通订单：
			//      非货到付款的，要已结算才能发货 OR 货到付款的，已确认就可以发货
			// 拼团订单：
			//      已经成团的
			str := " and (" + " ( payment_type!='cod' and (order_type='" + consts.OrderTypeNormal + "' or order_type='" +
				consts.OrderTypeChange + "' or order_type='" + consts.OrderTypeSupplyAgain + "')  and  order_status='" +
				consts.OrderStatusPaidOff + "')  " + " or ( payment_type='cod' and order_type='" + consts.OrderTypeNormal +
				"'  and  order_status='" + consts.OrderStatusConfirm + "') " + " or ( order_type='" + consts.OrderTypePinTuan +
				"'  and  order_status='" + consts.OrderStatusFormed + "') " + ")"
			sqlString.WriteString(str)
			break
		case consts.OrderTagWaitRog: //待收货
			str := " and o.order_status='" + consts.OrderStatusShipped + "'"
			sqlString.WriteString(str)
			break
		case consts.OrderTagWaitComment: //待评论
			str := " and o.ship_status='" + consts.ShipStatusShipRog + "' and o.comment_status='" + consts.CommentStatusUnfinished + "' "
			sqlString.WriteString(str)
			break
		case consts.OrderTagWaitChase: //待追评
			str := " and o.ship_status='" + consts.ShipStatusShipRog + "' and o.comment_status='" + consts.CommentStatusWaitChase + "' "
			sqlString.WriteString(str)
			break
		case consts.OrderTagCancelled: //已取消
			str := " and o.order_status='" + consts.OrderStatusCancelled + "'"
			sqlString.WriteString(str)
			break
		case consts.OrderTagComplete:
			str := " and o.order_status='" + consts.OrderStatusComplete + "'"
			sqlString.WriteString(str)
			break
		default:
			break
		}
	}

	sqlString.WriteString(" order by o.order_id desc")
	if pageNo != "" && pageSize != "" {
		pageNo, _ := strconv.Atoi(pageNo)
		pageSize, _ := strconv.Atoi(pageSize)
		sqlString.WriteString(sql_utils.LimitOffset(pageNo, pageSize))
	}

	rows := om.QuerySql(sqlString.String())
	defer rows.Close()

	orderList, _ := sql_utils.ParseJSON(rows)

	//订单自动取消天数
	cancelLeftDay := om.getCancelLeftDay()
	for _, order := range orderList {
		// TODO 待补充
		fmt.Println(cancelLeftDay, order)
	}

	return orderList, om.count()
}

func (om *OrderModel) count() (rows int64) {
	var (
		sql = "select count(1) from es_order o where disabled=0 "
	)

	err := om.QueryRow(sql).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}

	return rows
}

func (om *OrderModel) getCancelLeftDay() int {
	tradeSetting := CreateSettingFactory("").Get(consts.TRADE)

	if cancelOrderDay, ok := tradeSetting["cancel_order_day"].(int); ok {
		return cancelOrderDay
	}
	return 0
}

func (om *OrderModel) getOrder(orderId string) map[string]interface{} {
	rows := om.QuerySql("select * from es_order where sn = ?", orderId)
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
