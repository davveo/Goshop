package request

type CategoryRequest struct {
	Name          string `json:"name"`
	ParentId      string `json:"parent_id"`
	CategoryOrder string `json:"category_order"`
	Image         string `json:"image"`
	IsShow        string `json:"is_show"`
	AdvImage      string `json:"advImage"`
	AdvImageLink  string `json:"advImageLink"`
}
