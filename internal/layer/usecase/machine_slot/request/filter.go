package request

import (
	"vm-backend/pkg/db"
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
		PtrOrder(r.SortBy).
		AddWhere("machine_id = ?", r.MachineID).
		PtrWhere("id = ?", r.ID).
		PtrWhere("product_id = ?", r.ProductID).
		PtrWhere("code = ?", r.Code).
		PtrWhere("stock = ?", r.Stock).
		PtrWhere("capacity = ?", r.Capacity).
		PtrPreloads(r.Preloads)
}
