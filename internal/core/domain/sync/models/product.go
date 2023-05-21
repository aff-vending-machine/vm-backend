package models

import (
	"vm-backend/internal/core/domain/catalog"
)

type Product struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
}

func (m *Product) ToEntity() *catalog.Product {
	return &catalog.Product{
		SKU:          m.SKU,
		Name:         m.Name,
		ImageURL:     m.ImageURL,
		ProductPrice: m.Price,
		SalePrice:    m.Price,
	}
}
