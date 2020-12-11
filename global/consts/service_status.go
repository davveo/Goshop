package consts

const (
	/** 退货 */
	ServiceTypeReturnGoods = "RETURN_GOODS"

	/** 换货 */
	ServiceTypeChangeGoods = "CHANGE_GOODS"

	/** 补发商品 */
	ServiceTypeSupplyAgainGoods = "SUPPLY_AGAIN_GOODS"

	/** 取消订单 */
	ServiceTypeOrderCancel = "ORDER_CANCEL"
)

const (
	/** 待审核 */
	ServiceStatusApply = "APPLY"

	/** 审核通过 */
	ServiceStatusPass = "PASS"

	/** 审核未通过 */
	ServiceStatusRefuse = "REFUSE"

	/** 已退还商品 */
	ServiceStatusFullCourier = "FULL_COURIER"

	/** 待人工处理 */
	ServiceStatusWaitForManual = "WAIT_FOR_MANUAL"

	/** 已入库 */
	ServiceStatusStockIn = "STOCK_IN"

	/** 退款中 */
	ServiceStatusREFUNDING = "REFUNDING"

	/** 退款失败 */
	ServiceStatusREFUNDFAIL = "REFUNDFAIL"

	/** 已完成 */
	ServiceStatusCOMPLETED = "COMPLETED"

	/** 已关闭 */
	ServiceStatusCLOSED = "CLOSED"

	/** 异常状态 */
	ServicestatuserrorException = "ERROR_EXCEPTION"
)

var ServiceTypeMap = map[string]string{
	ServiceTypeReturnGoods:      "退货",
	ServiceTypeChangeGoods:      "换货",
	ServiceTypeSupplyAgainGoods: "补发商品",
	ServiceTypeOrderCancel:      "取消订单",
}

var ServiceStatusMap = map[string]string{
	ServiceStatusApply:          "待审核",
	ServiceStatusPass:           "审核通过",
	ServiceStatusRefuse:         "审核未通过",
	ServiceStatusFullCourier:    "已退还商品",
	ServiceStatusWaitForManual:  "待人工处理",
	ServiceStatusStockIn:        "已入库",
	ServiceStatusREFUNDING:      "退款中",
	ServiceStatusREFUNDFAIL:     "退款失败",
	ServiceStatusCOMPLETED:      "已完成",
	ServiceStatusCLOSED:         "已关闭",
	ServicestatuserrorException: "异常状态",
}
