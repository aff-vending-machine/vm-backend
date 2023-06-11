package request

import (
	"vm-backend/pkg/helpers/db"
)

type Delete struct {
	ID       uint  `json:"id" query:"id" validate:"required"`
	BranchID *uint `json:"branch_id,omitempty" query:"branch_id"`
}

func (r *Delete) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID).
		PtrWhere("branch_id = ?", r.BranchID)
}
