package request

import (
	"vm-backend/pkg/helpers/conv"
	"vm-backend/pkg/helpers/db"
)

type Update struct {
	ID           uint     `json:"id" query:"id" validate:"required"`
	GroupID      *uint    `json:"group_id,omitempty"`
	Name         *string  `json:"name,omitempty"`
	ImageURL     *string  `json:"image_url,omitempty"`
	ProductPrice *float64 `json:"product_price,omitempty"`
	SalePrice    *float64 `json:"sale_price,omitempty"`
	IsEnable     *bool    `json:"is_enable,omitempty"`
}

func (r *Update) ToQuery() *db.Query {
	return db.NewQuery().
		Where("id = ?", r.ID)
}

func (r *Update) ToUpdate() map[string]interface{} {
	result, _ := conv.ToMap(r)
	delete(result, "id")
	return result
}
