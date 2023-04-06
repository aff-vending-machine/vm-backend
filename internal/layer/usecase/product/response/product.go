package response

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Product struct {
	ID       uint    `json:"id"`
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
}

func ProductEntityToView(e *entity.Product) *Product {
	return &Product{
		ID:       e.ID,
		SKU:      e.SKU,
		Name:     e.Name,
		Type:     e.Type,
		ImageURL: e.ImageURL,
		Price:    e.Price,
	}
}

func ProductEntityToList(es []entity.Product) []Product {
	items := make([]Product, len(es))
	for i, e := range es {
		items[i] = *ProductEntityToView(&e)
	}

	return items
}
