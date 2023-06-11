package request

import (
	"vm-backend/pkg/helpers/db"
)

type Filter struct {
	Limit    *int    `json:"limit,omitempty" query:"limit"`
	Offset   *int    `json:"offset,omitempty" query:"offset"`
	SortBy   *string `json:"sort_by,omitempty" query:"sort_by"`
	Preloads *string `json:"preloads,omitempty" query:"preloads"`
	ID       *uint   `json:"id,omitempty" query:"id"`
	Username *string `json:"username,omitempty" query:"username"`
}

func (r *Filter) ToQuery() *db.Query {
	return db.NewQuery().
		PtrLimit(r.Limit).
		PtrOffset(r.Offset).
		PtrOrder(r.SortBy).
		PtrWhere("id = ?", r.ID).
		PtrWhere("username = ?", r.Username).
		PtrPreloads(r.Preloads)
}
