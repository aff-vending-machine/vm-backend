package response

type Product struct {
	ID           uint    `json:"id"`
	GroupID      uint    `json:"group_id"`
	Group        *Group  `json:"group,omitempty"`
	SKU          string  `json:"sku"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ImageURL     string  `json:"image_url"`
	Barcode      string  `json:"barcode"`
	ProductPrice float64 `json:"product_price"`
	SalePrice    float64 `json:"sale_price"`
	IsEnable     bool    `json:"is_enable"`
}
