package request

import (
	"vm-backend/pkg/db"
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
		PtrLimit(r.Limit).
		PtrOffset(r.Offset).
		PtrOrder(r.SortBy).
		PtrWhere("id = ?", r.ID).
		PtrWhere("branch_id = ?", r.BranchID).
		PtrWhere("name = ?", r.Name).
		PtrWhere("serial_number = ?", r.SerialNumber).
		PtrWhere("status = ?", r.Status).
		PtrPreloads(r.Preloads)
}
