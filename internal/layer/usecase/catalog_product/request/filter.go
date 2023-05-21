package request

import "vm-backend/pkg/db"

type Filter struct {
	Limit    *int    `json:"limit,omitempty" query:"limit"`
	Offset   *int    `json:"offset,omitempty" query:"offset"`
	SortBy   *string `json:"sort_by,omitempty" query:"sort_by"`
	Preloads *string `json:"preloads,omitempty" query:"preloads"`
	GroupID  *uint   `json:"group_id,omitempty" query:"group_id"`
	ID       *uint   `json:"id,omitempty" query:"id"`
	Name     *string `json:"name,omitempty" query:"name"`
	SKU      *string `json:"sku,omitempty" query:"sku"`
}

func (r *Filter) ToQuery() *db.Query {
	return db.NewQuery().
		PtrOffset(r.Offset).
		PtrLimit(r.Limit).
		PtrOrder(r.SortBy).
		PtrWhere("group_id = ?", r.GroupID).
		PtrWhere("id = ?", r.ID).
		PtrWhere("name = ?", r.Name).
		PtrWhere("sku = ?", r.SKU).
		PtrPreloads(r.Preloads)
}
