package consts

const (
	/** 商家审核操作 */
	ServiceOperateTypeSellerAudit = "SELLER_AUDIT"

	/** 商家入库操作 */
	ServiceOperateTypeStockIn = "STOCK_IN"

	/** 商家退款操作 */
	ServiceOperateTypeSellerRefund = "SELLER_REFUND"

	/** 平台退款操作 */
	ServiceOperateTypeAdminRefund = "ADMIN_REFUND"

	/** 买家填写物流信息操作 */
	ServiceOperateTypeFillLogisticsInfo = "FILL_LOGISTICS_INFO"

	/** 关闭操作 */
	ServiceOperateTypeClose = "CLOSE"

	/** 创建新订单操作 */
	ServiceOperateTypeCreateNewOrder = "CREATE_NEW_ORDER"
)

var ServiceOperateTypeMap = map[string]string{
	ServiceOperateTypeSellerAudit:       "商家审核操作",
	ServiceOperateTypeStockIn:           "审核入库操作",
	ServiceOperateTypeSellerRefund:      "商家退款操作",
	ServiceOperateTypeAdminRefund:       "平台退款操作",
	ServiceOperateTypeFillLogisticsInfo: "买家填写物流信息操作",
	ServiceOperateTypeClose:             "关闭操作",
	ServiceOperateTypeCreateNewOrder:    "创建新订单操作",
}
