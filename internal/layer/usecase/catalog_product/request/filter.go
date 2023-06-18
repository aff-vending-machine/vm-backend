package request

import "vm-backend/pkg/helpers/db"

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
		OffsetIf(r.Offset).
		LimitIfNotNil(r.Limit).
		OrderIf(r.SortBy).
		WhereIf("group_id = ?", r.GroupID).
		WhereIf("id = ?", r.ID).
		WhereIf("name = ?", r.Name).
		WhereIf("sku = ?", r.SKU).
		PreloadsIf(r.Preloads)
}
