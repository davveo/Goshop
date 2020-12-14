package consts

const (
	RefundWayOriginal = "ORIGINAL"
	RefundWayOffline  = "OFFLINE"
	RefundWayAccount  = "ACCOUNT"
)

var RefundWayMap = map[string]string{
	RefundWayOriginal: "原路退回",
	RefundWayOffline:  "线下退款",
	RefundWayAccount:  "账户退款",
}
