package request

import (
	"vm-backend/pkg/helpers/db"
)

type Filter struct {
	Limit        *int    `json:"limit,omitempty" query:"limit"`
	Offset       *int    `json:"offset,omitempty" query:"offset"`
	SortBy       *string `json:"sort_by,omitempty" query:"sort_by"`
	Preloads     *string `json:"preloads,omitempty" query:"preloads"`
	ID           *uint   `json:"id,omitempty" query:"id"`
	BranchID     *uint   `json:"branch_id,omitempty" query:"branch_id"`
	Name         *string `json:"name,omitempty" query:"name"`
	SerialNumber *string `json:"serial_number,omitempty" query:"serial_number"`
	Status       *bool   `json:"status,omitempty" query:"status"`
}

func (r *Filter) ToQuery() *db.Query {
	return db.NewQuery().
		LimitIfNotNil(r.Limit).
		OffsetIf(r.Offset).
		OrderIf(r.SortBy).
		WhereIf("id = ?", r.ID).
		WhereIf("branch_id = ?", r.BranchID).
		WhereIf("name = ?", r.Name).
		WhereIf("serial_number = ?", r.SerialNumber).
		WhereIf("status = ?", r.Status).
		PreloadsIf(r.Preloads)
}
