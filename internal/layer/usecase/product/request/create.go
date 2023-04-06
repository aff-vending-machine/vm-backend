package request

import "github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"

type Create struct {
	SKU      string  `json:"sku" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Type     string  `json:"type" validate:"required"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price" validate:"required"`
}

func (r *Create) ToEntity() *entity.Product {
	return &entity.Product{
		SKU:      r.SKU,
		Name:     r.Name,
		Type:     r.Type,
		ImageURL: r.ImageURL,
		Price:    r.Price,
	}
}
