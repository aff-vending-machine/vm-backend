package model

type Product struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
}
