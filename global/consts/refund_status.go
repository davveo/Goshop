package consts

const (
	RefundStatusApply      = "APPLY"
	RefundStatusRefunding  = "REFUNDING"
	RefundStatusRefundFail = "REFUNDFAIL"
	RefundStatusCompleted  = "COMPLETED"
)

var RefundStatusMap = map[string]string{
	RefundStatusApply:      "待退款",
	RefundStatusRefunding:  "退款中",
	RefundStatusRefundFail: "退款失败",
	RefundStatusCompleted:  "完成",
}
