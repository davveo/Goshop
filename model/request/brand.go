package request

type BrandRequest struct {
	Name     string `json:"name"`
	Logo     string `json:"logo"`
	BrandId  string `json:"brand_id"`
	Disabled string `json:"disabled"`
}
