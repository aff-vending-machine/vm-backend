package request

import "vm-backend/pkg/helpers/db"

type Filter struct {
	Limit    *int    `json:"limit,omitempty" query:"limit"`
	Offset   *int    `json:"offset,omitempty" query:"offset"`
	SortBy   *string `json:"sort_by,omitempty" query:"sort_by"`
	Preloads *string `json:"preloads,omitempty" query:"preloads"`
	ID       *uint   `json:"id,omitempty" query:"id"`
	Channel  *string `json:"channel,omitempty" query:"channel"`
	Vendor   *string `json:"vendor,omitempty" query:"vendor"`
	IsEnable *bool   `json:"is_enable,omitempty" query:"is_enable"`
}

func (r *Filter) ToQuery() *db.Query {
	return db.NewQuery().
		LimitIfNotNil(r.Limit).
		OffsetIf(r.Offset).
		OrderIf(r.SortBy).
		WhereIf("id = ?", r.ID).
		WhereIf("channel = ?", r.Channel).
		WhereIf("vendor = ?", r.Vendor).
		WhereIf("is_enable = ?", r.IsEnable).
		PreloadsIf(r.Preloads)
}
