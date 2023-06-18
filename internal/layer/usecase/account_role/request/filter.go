package request

import "vm-backend/pkg/helpers/db"

type Filter struct {
	Limit    *int    `json:"limit,omitempty" query:"limit"`
	Offset   *int    `json:"offset,omitempty" query:"offset"`
	SortBy   *string `json:"sort_by,omitempty" query:"sort_by"`
	Preloads *string `json:"preloads,omitempty" query:"preloads"`
	ID       *uint   `json:"id,omitempty" query:"id"`
	Name     *string `json:"name,omitempty" query:"name"`
}

func (r *Filter) ToQuery() *db.Query {
	return db.NewQuery().
		OffsetIf(r.Offset).
		LimitIfNotNil(r.Limit).
		OrderIf(r.SortBy).
		WhereIf("id = ?", r.ID).
		WhereIf("name = ?", r.Name).
		PreloadsIf(r.Preloads)
}
