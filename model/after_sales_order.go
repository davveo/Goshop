package model

import (
	"Goshop/global/consts"
	"Goshop/utils/yml_config"
	"bytes"
	"errors"
	"fmt"
	"log"
)

func CreateAfterSaleOrderFactory(sqlType string) *AfterSaleOrderModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType")
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &AfterSaleOrderModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("AfterSaleOrderModel工厂初始化失败")
	return nil
}

type AfterSaleOrderModel struct {
	*BaseModel
}

func (asom *AfterSaleOrderModel) updateServiceStatus(serviceSn, serviceStatus, auditRemark,
	storageRemark, refundRemark, closeReason string) error {
	var sqlString bytes.Buffer
	sqlString.WriteString(fmt.Sprintf("update es_as_order set service_status = %s", serviceStatus))

	if auditRemark != "" {
		sqlString.WriteString(fmt.Sprintf(",audit_remark = %s", auditRemark))
	}
	if storageRemark != "" {
		sqlString.WriteString(fmt.Sprintf(",stock_remark = %s", storageRemark))
	}
	if refundRemark != "" {
		sqlString.WriteString(fmt.Sprintf(",refund_remark = %s", refundRemark))
	}

	if closeReason != "" {
		sqlString.WriteString(fmt.Sprintf(",close_reason = %s", closeReason))
	}

	sqlString.WriteString(fmt.Sprintf(" where sn = %s", serviceSn))

	if asom.ExecuteSql(sqlString.String()) == -1 {
		return errors.New("删除售后订单失败")
	}

	return nil
}

func (asom *AfterSaleOrderModel) setServiceNewOrderSn(serviceSn, newOrderSn string) error {
	if asom.ExecuteSql("update es_as_order set service_status = ?,new_order_sn = ? where sn = ?",
		consts.ServiceStatusPass, newOrderSn, serviceSn) == -1 {
		return errors.New("更新售后订单失败")
	}
	return nil
}

func (asom *AfterSaleOrderModel) editAfterSaleShopName(shopId, shopName string) error {
	if asom.ExecuteSql("update es_as_order set seller_name = ? where seller_id = ? ", shopName, shopId) == -1 {
		return errors.New("更新售后店铺名失败")
	}
	return nil
}
