package request

type BatchAuditRequest struct {
	GoodsIds []string `json:"goods_ids"`
	Message  string   `json:"message"`
	Pass     int      `json:"pass"`
}
