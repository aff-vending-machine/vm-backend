package request

import "vm-backend/pkg/helpers/db"

type Get struct {
	ID       uint    `json:"id" query:"id" validate:"required"`
	BranchID *uint   `json:"branch_id" query:"branch_id"`
	Preloads *string `json:"preloads,omitempty" query:"preloads"`
}

func (r *Get) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID).
		PtrWhere("branch_id = ?", r.BranchID).
		AddPreload(*r.Preloads)
}
