package request

import "vm-backend/pkg/helpers/db"

type Get struct {
	ID       uint    `json:"id" query:"id" validate:"required"`
	BranchID *uint   `json:"branch_id,omitempty" query:"branch_id"`
	Preloads *string `json:"preloads,omitempty" query:"preloads"`
}

func (r *Get) ToQuery() *db.Query {
	return db.NewQuery().
		Where("id = ?", r.ID).
		WhereIf("branch_id = ?", r.BranchID).
		PreloadsIf(r.Preloads)
}
