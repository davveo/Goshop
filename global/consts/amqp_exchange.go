package consts

// AmqpExchange
const (
	ExchangeTestExchange = "TEST_EXCHANGE_FANOUT"

	/**
	 * PC首页变化消息
	 */
	ExchangePcIndexChange = "PC_INDEX_CHANGE"

	/**
	 * 移动端首页变化消息
	 */
	ExchangeMobileIndexChange = "MOBILE_INDEX_CHANGE"

	/**
	 * 商品变化消息
	 */
	ExchangeGoodsChange = "GOODS_CHANGE"
	/**
	 * 商品优先级变化消息
	 */
	ExchangeGoodsPriorityChange = "GOODS_PRIORITY_CHANGE"

	/**
	 * 商品变化消息附带原因
	 */
	ExchangeGoodsChangeReason = "GOODS_CHANGE_REASON"

	/**
	 * 帮助变化消息
	 */
	ExchangeHelpChange = "HELP_CHANGE"

	/**
	 * 页面生成消息
	 */
	ExchangePageCreate = "PAGE_CREATE"

	/**
	 * 索引生成消息
	 */
	ExchangeIndexCreate = "INDEX_CREATE"

	/**
	 * 订单创建消息
	 * 没有入库
	 */
	ExchangeOrderCreate = "ORDER_CREATE"

	/**
	 * 入库失败消息
	 * 入库失败
	 */
	ExchangeOrderIntodbError = "ORDER_INTODB_ERROR"

	/**
	 * 订单状态变化消息
	 * 带入库的
	 */
	ExchangeOrderStatusChange = "ORDER_STATUS_CHANGE"

	/**
	 * 会员登录消息
	 */
	ExchangeMemeberLogin = "MEMEBER_LOGIN"

	/**
	 * 会员注册消息
	 */
	ExchangeMemeberRegister = "MEMEBER_REGISTER"

	/**
	 * 店铺变更消息
	 */
	ExchangeShopChangeRegister = "SHOP_CHANGE_REGISTER"
	/**
	 * 分类变更消息
	 */
	ExchangeGoodsCategoryChange = "GOODS_CATEGORY_CHANGE"

	/**
	 * 售后状态改变消息
	 */
	ExchangeRefundStatusChange = "REFUND_STATUS_CHANGE"

	/**
	 * 发送站内信息
	 */
	ExchangeMemberMessage = "MEMBER_MESSAGE"

	/**
	 * 发送手机短信消息
	 */
	ExchangeSendMessage = "SEND_MESSAGE"

	/**
	 * 邮件发送消息
	 */
	ExchangeEmailSendMessage = "EMAIL_SEND_MESSAGE"

	/**
	 * 商品评论消息
	 */
	ExchangeGoodsCommentComplete = "GOODS_COMMENT_COMPLETE"
	/**
	 * 网上支付
	 */
	ExchangeOnlinePay = "ONLINE_PAY"

	/**
	 * 完善个人资料
	 */
	ExchangeMemberInfoComplete = "MEMBER_INFO_COMPLETE"

	/**
	 * 站点导航栏变化消息
	 */
	ExchangeSiteNavigationChange = "SITE_NAVIGATION_CHANGE"

	/**
	 * 商品收藏
	 */
	ExchangeGoodsCollectionChange = "GOODS_COLLECTION_CHANGE"

	/**
	 * 店铺收藏
	 */
	ExchangeSellerCollectionChange = "SELLER_COLLECTION_CHANGE"

	/**
	 * 店铺关闭
	 */
	ExchangeCloseStore = "CLOSE_STORE"

	/**
	 * 店铺开启
	 */
	ExchangeOpenStore = "OPEN_STORE"

	/**
	 * 店铺信息发生改变
	 */
	ExchangeShopChange = "SHOP_CHANGE"

	/**
	 * 店铺浏览统计
	 */
	ExchangeShopViewCount = "SHOP_VIEW_COUNT"

	/**
	 * 商品浏览统计
	 */
	ExchangeGoodsViewCount = "GOODS_VIEW_COUNT"

	/**
	 * 会员资料改变
	 */
	ExchangeMemberInfoChange = "MEMBER_INFO_CHANGE"

	/**
	 * 会员历史足迹
	 */
	ExchangeMemberHistory = "MEMBER_HISTORY"

	/**
	 * 搜索关键字消息
	 */
	ExchangeSearchKeywords = "SEARCH_KEYWORDS"

	/**
	 * 提示词变更
	 */
	ExchangeGoodsWordsChange = "GOODS_WORDS_CHANGE"

	/**
	 * 拼团成功消息
	 */
	ExchangePintuanSuccess = "PINTUAN_SUCCESS"

	/**
	 * 运费模板变化消息
	 */
	ExchangeShipTemplateChange = "SHIP_TEMPLATE_CHANGE"
	/**
	 * 会员商品咨询消息
	 */
	ExchangeMemberGoodsAsk = "MEMBER_GOODS_ASK"

	/**
	 * 会员商品咨询回复消息
	 */
	ExchangeMemberGoodsAskReply = "MEMBER_GOODS_ASK_REPLY"

	/**
	 * 商家创建换货或补发商品售后服务新订单消息
	 */
	ExchangeAsSellerCreateOrder = "AS_SELLER_CREATE_ORDER"

	/**
	 * 售后服务单状态改变消息
	 */
	ExchangeAsStatusChange = "AS_STATUS_CHANGE"
)
