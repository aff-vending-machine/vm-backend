package request

import "vm-backend/pkg/helpers/db"

type ResetPassword struct {
	ID       uint  `json:"id" query:"id" validate:"required"`
	Level    int   `json:"level" query:"level" validate:"required"`
	BranchID *uint `json:"branch_id,omitempty" query:"branch_id"`
}

func (r *ResetPassword) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID).
		PtrWhere("branch_id = ?", r.BranchID)
}

func (r *ResetPassword) ToUpdate(hashed string) map[string]interface{} {
	return map[string]interface{}{
		"password": hashed,
	}
}
