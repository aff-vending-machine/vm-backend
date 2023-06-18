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
	RoleID   *uint   `json:"role_id,omitempty" query:"role_id"`
	BranchID *uint   `json:"branch_id,omitempty" query:"branch_id"`
	Username *string `json:"username,omitempty" query:"username"`
}

func (r *Filter) ToQuery() *db.Query {
	return db.NewQuery().
		LimitIfNotNil(r.Limit).
		OffsetIf(r.Offset).
		OrderIf(r.SortBy).
		WhereIf("id = ?", r.ID).
		WhereIf("role_id = ?", r.RoleID).
		WhereIf("branch_id = ?", r.BranchID).
		WhereIf("username = ?", r.Username).
		PreloadsIf(r.Preloads)
}
