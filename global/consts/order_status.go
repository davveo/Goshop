package consts

// OrderStatus
const (
	/**
	 * 新订单
	 */
	OrderStatusNew = "NEW"

	/**
	 * 出库失败
	 */
	OrderStatusIntodbError = "INTODB_ERROR"

	/**
	 * 已确认
	 */
	OrderStatusConfirm = "CONFIRM"

	/**
	 * 已付款
	 */
	OrderStatusPaidOff = "PAID_OFF"

	/**
	 * 已成团
	 */
	OrderStatusFormed = "FORMED"

	/**
	 * 已发货
	 */
	OrderStatusShipped = "SHIPPED"

	/**
	 * 已收货
	 */
	OrderStatusRog = "ROG"

	/**
	 * 已完成
	 */
	OrderStatusComplete = "COMPLETE"

	/**
	 * 已取消
	 */
	OrderStatusCancelled = "CANCELLED"

	/**
	 * 售后中
	 */
	OrderStatusAfterService = "AFTER_SERVICE"
)

// OrderTag
const (
	// 所有订单
	OrderTagAll = "ALL"
	/** 待付款 */
	OrderTagWaitPay = "WAIT_PAY"

	/** 待发货 */
	OrderTagWaitShip = "WAIT_SHIP"

	/** 待收货 */
	OrderTagWaitRog = "WAIT_ROG"

	/** 已取消 */
	OrderTagCancelled = "CANCELLED"

	/** 已完成 */
	OrderTagComplete = "COMPLETE"

	/** 待评论 */
	OrderTagWaitComment = "WAIT_COMMENT"

	/**
	 * 待追评
	 */
	OrderTagWaitChase = "WAIT_CHASE"

	/** 售后中 */
	OrderTagRefund = "REFUND"
)

// 订单来源
const (
	// pc客户端
	ClientTypePc = "PC"
	// WAP:WAP客户端
	ClientTypeWap = "WAP"
	// NATIVE:原生APP
	ClientTypeNative = "NATIVE"
	// REACT:RNAPP
	ClientTypeReact = "REACT"
	// MINI:小程序
	ClientTypeMini = "MINI"
)

// 支付方式
const (
	// ONLINE:在线支付
	PayOnline = "ONLINE"
	// COD:货到付款
	PayCod = "COD"
)

// OrderType
const (
	OrderTypeNormal      = "NORMAL"       // 普通订单
	OrderTypePinTuan     = "PINTUAN"      // 拼团订单
	OrderTypeChange      = "CHANGE"       // 换货订单
	OrderTypeSupplyAgain = "SUPPLY_AGAIN" // 补发商品订单
)

// 发货状态
const (

	/**
	 * 未发货
	 */
	ShipStatusShipNo = "SHIP_NO"
	/**
	 * 已发货
	 */
	ShipStatusShipYes = "SHIP_YES"
	/**
	 * 已收货
	 */
	ShipStatusShipRog = "SHIP_ROG"
)
