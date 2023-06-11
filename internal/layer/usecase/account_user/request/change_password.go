package request

import (
	"vm-backend/pkg/helpers/db"
)

type ChangePassword struct {
	ID          uint   `json:"id" query:"id" validate:"required"`
	BranchID    *uint  `json:"branch_id,omitempty" query:"branch_id"`
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

func (r *ChangePassword) ToQuery() *db.Query {
	return db.NewQuery().
	AddWhere("id = ?", r.ID).
	PtrWhere("branch_id = ?", r.BranchID)
}

func (r *ChangePassword) ToUpdate(hashed string) map[string]interface{} {
	return map[string]interface{}{
		"password": hashed,
	}
}
