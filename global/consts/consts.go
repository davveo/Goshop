package consts

// 这里定义的常量，一般是具有错误代码+错误说明组成，一般用于接口返回
const (
	// 进程被结束
	ProcessKilled string = "收到信号，进程被结束"
	// 表单验证器前缀
	ValidatorPrefix              string = "Form_Validator_"
	ValidatorParamsCheckFailCode int    = -400300
	ValidatorParamsCheckFailMsg  string = "参数校验失败"

	//服务器代码发生错误
	ServerOccurredErrorCode int    = -500100
	ServerOccurredErrorMsg  string = "服务器内部发生代码执行错误,检测文件mime类型发生错误。"

	// token相关
	JwtTokenSignKey         string = "github.com/davvo/goshop"
	JwtTokenCreatedExpireAt int64  = 3600    // 创建时token默认有效期3600秒
	JwtTokenRefreshExpireAt int64  = 7200    // 刷新token时，延长7200秒
	JwtTokenOK              int    = 200100  //token有效
	JwtTokenInvalid         int    = -400100 //无效的token
	JwtTokenExpired         int    = -400101 //过期的token
	JwtTokenOnlineUsers     int    = 10      // 设置一个账号最大允许几个用户同时在线，默认为10

	//snowflake错误
	SnowFlakeMachineId      int16  = 1024
	SnowFlakeMachineIllegal string = "SnowFlake数据越界，大于65535"

	// CURD 常用业务状态码
	CurdStatusOkCode         int    = 200
	CurdStatusOkMsg          string = "Success"
	CurdCreatFailCode        int    = -400200
	CurdCreatFailMsg         string = "新增失败"
	CurdUpdateFailCode       int    = -400201
	CurdUpdateFailMsg        string = "更新失败"
	CurdDeleteFailCode       int    = -400202
	CurdDeleteFailMsg        string = "删除失败"
	CurdSelectFailCode       int    = -400203
	CurdSelectFailMsg        string = "查询无数据"
	CurdRegisterFailCode     int    = -400204
	CurdRegisterFailMsg      string = "注册失败"
	CurdLoginFailCode        int    = -400205
	CurdLoginFailMsg         string = "登录失败"
	CurdRefreshTokenFailCode int    = -400206
	CurdRefreshTokenFailMsg  string = "刷新Token失败"

	//文件上传
	FilesUploadFailCode            int    = -400250
	FilesUploadFailMsg             string = "文件上传失败, 获取上传文件发生错误!"
	FilesUploadMoreThanMaxSizeCode int    = -400251
	FilesUploadMoreThanMaxSizeMsg  string = "长传文件超过系统设定的最大值,系统允许的最大值（M）："
	FilesUploadMimeTypeFailCode    int    = -400252
	FilesUploadMimeTypeFailMsg     string = "文件mime类型不允许"

	//websocket
	WsServerNotStartCode int    = -400300
	WsServerNotStartMsg  string = "websocket 服务没有开启，请在配置文件开启，相关路径：config/config.yml"
	WsOpenFailCode       int    = -400301
	WsOpenFailMsg        string = "websocket open阶段初始化基本参数失败"

	//验证码
	CaptchaGetParamsInvalidMsg    string = "获取验证码：提交的验证码参数无效,请检查验证码ID以及文件名后缀是否完整"
	CaptchaGetParamsInvalidCode   int    = -400350
	CaptchaCheckParamsInvalidMsg  string = "校验验证码：提交的参数无效，请确保提交的验证码ID和值有效"
	CaptchaCheckParamsInvalidCode int    = -400351
	CaptchaCheckOkMsg             string = "验证码校验通过"
	CaptchaCheckOkCode            int    = 200
	CaptchaCheckFailCode          int    = -400355
	CaptchaCheckFailMsg           string = "验证码校验失败"

	// 系统设置参数
	SYSTEM       = "SYSTEM"
	SITE         = "SITE"
	PHOTO_SIZE   = "PHOTO_SIZE"
	POINT        = "POINT"
	DISTRIBUTION = "DISTRIBUTION"
	TEST         = "TEST"
	PUSH         = "PUSH"
	PAGE         = "PAGE"
	ES_SIGN      = "ES_SIGN"

	// 其他常数
	YEAR  = "YEAR"
	MONTH = "MONTH"

	// cache prefix
	NONCE                      = "NONCE"
	TOKEN                      = "TOKEN"
	SETTING                    = "SETTING"
	EXPRESS                    = "EXPRESS"
	CAPTCHA                    = "CAPTCHA"
	GOODS                      = "GOODS"
	SHIP_SCRIPT                = "SHIP_SCRIPT"
	SKU                        = "SKU"
	SKU_STOCK                  = "SKU_STOCK"
	GOODS_STOCK                = "GOODS_STOCK"
	GOODS_CAT                  = "GOODS_CAT"
	VISIT_COUNT                = "VISIT_COUNT"
	UPLOADER                   = "UPLOADER"
	REGION                     = "REGION"
	SPlATFORM                  = "SPlATFORM"
	_CODE_PREFIX               = "_CODE_PREFIX"
	SMTP                       = "SMTP"
	SETTINGS                   = "SETTINGS"
	WAYBILL                    = "WAYBILL"
	SMS_CODE                   = "SMS_CODE"
	EMAIL_CODE                 = "EMAIL_CODE"
	ADMIN_URL_ROLE             = "ADMIN_URL_ROLE"
	SHOP_URL_ROLE              = "SHOP_URL_ROLE"
	MOBILE_VALIDATE            = "MOBILE_VALIDATE"
	EMAIL_VALIDATE             = "EMAIL_VALIDATE"
	SHIP_TEMPLATE              = "SHIP_TEMPLATE"
	PROMOTION_KEY              = "PROMOTION_KEY"
	SHIP_TEMPLATE_ONE          = "SHIP_TEMPLATE_ONE"
	STORE_ID_MINUS_KEY         = "STORE_ID_MINUS_KEY"
	STORE_ID_HALF_PRICE_KEY    = "STORE_ID_HALF_PRICE_KEY"
	STORE_ID_FULL_DISCOUNT_KEY = "STORE_ID_FULL_DISCOUNT_KEY"
	STORE_ID_SECKILL_KEY       = "STORE_ID_SECKILL_KEY"
	STORE_ID_GROUP_BUY_KEY     = "STORE_ID_GROUP_BUY_KEY"
	STORE_ID_EXCHANGE_KEY      = "STORE_ID_EXCHANGE_KEY"
	CART_ORIGIN_DATA_PREFIX    = "CART_ORIGIN_DATA_PREFIX"
	BUY_NOW_ORIGIN_DATA_PREFIX = "BUY_NOW_ORIGIN_DATA_PREFIX"
	TRADE_ORIGIN_DATA_PREFIX   = "TRADE_ORIGIN_DATA_PREFIX"
	CART_SKU_PREFIX            = "CART_SKU_PREFIX"
	CART_MEMBER_ID_PREFIX      = "CART_MEMBER_ID_PREFIX"
	CART_PROMOTION_PREFIX      = "CART_PROMOTION_PREFIX"
	PRICE_SESSION_ID_PREFIX    = "PRICE_SESSION_ID_PREFIX"
	TRADE_SESSION_ID_PREFIX    = "TRADE_SESSION_ID_PREFIX"
	CHECKOUT_PARAM_ID_PREFIX   = "CHECKOUT_PARAM_ID_PREFIX"
	TRADE_SN_CACHE_PREFIX      = "TRADE_SN_CACHE_PREFIX"
	ORDER_SN_CACHE_PREFIX      = "ORDER_SN_CACHE_PREFIX"
	ORDER_SN_SIGN_CACHE_PREFIX = "ORDER_SN_SIGN_CACHE_PREFIX"
	PAY_LOG_SN_CACHE_PREFIX    = "PAY_LOG_SN_CACHE_PREFIX"
	SMALL_CHANGE_CACHE_PREFIX  = "SMALL_CHANGE_CACHE_PREFIX"
	AFTER_SALE_SERVICE_PREFIX  = "AFTER_SALE_SERVICE_PREFIX"
	TRADE                      = "TRADE"
	GOODS_GRADE                = "GOODS_GRADE"
	REGIONALL                  = "REGIONALL"
	REGIONLIDEPTH              = "REGIONLIDEPTH"
	SITE_NAVIGATION            = "SITE_NAVIGATION"
	CONNECT_LOGIN              = "CONNECT_LOGIN"
	SESSION_KEY                = "SESSION_KEY"
	PAYMENT_CONFIG             = "PAYMENT_CONFIG"
	HOT_KEYWORD                = "HOT_KEYWORD"
	FLOW                       = "FLOW"
)

const (
	TimeFormatStyleV1 = "2006/01/02 15:04:05"
	TimeFormatStyleV2 = "20060102"
)

// 审核状态
const (
	WAIT_AUDIT   = "WAIT_AUDIT"   // 待审核
	PASS_AUDIT   = "PASS_AUDIT"   // 审核通过
	REFUSE_AUDIT = "REFUSE_AUDIT" // 审核拒绝
)

// 删除状态枚举
const (
	DELETED = "DELETED" // 已经删除
	NORMAL  = "NORMAL"  // 正常
)

// Permission
const (
	PermissionSELLER = iota
	PermissionBUYER
	PermissionADMIN
	PermissionCLIENT
)

// 商家操作枚举
const (
	/**
	 * 下架
	 */
	GoodsOperateUNDER = iota
	/**
	 * 还原
	 */
	GoodsOperateREVRET
	/**
	 * 放入回收站
	 */
	GoodsOperateRECYCLE
	/**
	 * 删除
	 */
	GoodsOperateDELETE
)

// 促销活动工具枚举
const (
	/**
	 * 不参与活动（指不参与任何单品活动）
	 */
	PromotionNo = "NO"

	/**
	 * 积分商品(积分活动)
	 */
	PromotionPoint = "POINT"

	/**
	 * 单品立减活动
	 */
	PromotionMinus = "MINUS"

	/**
	 * 团购活动
	 */
	PromotionGroupBuy = "GROUPBUY"

	/**
	 *积分换购活动
	 */
	PromotionExchange = "EXCHANGE"

	/**
	 * 第二件半价活动
	 */
	PromotionHalfPrice = "HALF_PRICE"

	/**
	 *满优惠活动
	 */
	PromotionFullDiscount = "FULL_DISCOUNT"

	/**
	 * 限时抢购
	 */
	PromotionSeckill = "SECKILL"

	/**
	 * 拼团活动类型
	 * 指定商品，优惠价格
	 */
	PromotionPinTuan = "PINTUAN"
)
