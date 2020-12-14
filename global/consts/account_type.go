package consts

const (
	AccountTypeAliPay       = "ALIPAY"
	AccountTypeWeiXinPay    = "WEIXINPAY"
	AccountTypeBankTransfer = "BANKTRANSFER"
)

var AccountTypeMap = map[string]string{
	AccountTypeAliPay:       "支付宝",
	AccountTypeWeiXinPay:    "微信",
	AccountTypeBankTransfer: "银行转账",
}
