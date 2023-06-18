package request

import (
	"vm-backend/pkg/helpers/db"
)

type Filter struct {
	MachineID uint    `json:"machine_id" query:"machine_id"`
	SortBy    *string `json:"sort_by,omitempty" query:"sort_by"`
	Preloads  *string `json:"preloads,omitempty" query:"preloads"`
	ID        *uint   `json:"id,omitempty" query:"id"`
	ProductID *uint   `json:"product_id,omitempty" query:"product_id"`
	Code      *string `json:"code,omitempty" query:"code"`
	Stock     *int    `json:"stock,omitempty" query:"stock"`
	Capacity  *int    `json:"capacity,omitempty" query:"capacity"`
}

func (r *Filter) ToQuery() *db.Query {
	return db.NewQuery().
		OrderIf(r.SortBy).
		Where("machine_id = ?", r.MachineID).
		WhereIf("id = ?", r.ID).
		WhereIf("product_id = ?", r.ProductID).
		WhereIf("code = ?", r.Code).
		WhereIf("stock = ?", r.Stock).
		WhereIf("capacity = ?", r.Capacity).
		PreloadsIf(r.Preloads)
}
