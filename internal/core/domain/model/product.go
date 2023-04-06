package model

import "github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"

type Product struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
}

func ToProduct(e *entity.Product) *Product {
	if e == nil {
		return nil
	}

	return &Product{
		SKU:      e.SKU,
		Name:     e.Name,
		Type:     e.Type,
		ImageURL: e.ImageURL,
		Price:    e.Price,
	}
}
