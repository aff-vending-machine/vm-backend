package request

type Create struct {
	GroupID      uint    `json:"group_id" validate:"required"`
	SKU          string  `json:"sku" validate:"required"`
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description"`
	ImageURL     string  `json:"image_url"`
	Barcode      string  `json:"barcode"`
	ProductPrice float64 `json:"product_price" validate:"required"`
	SalePrice    float64 `json:"sale_price" validate:"required"`
	IsEnable     bool    `json:"is_enable" validate:"required"`
}
