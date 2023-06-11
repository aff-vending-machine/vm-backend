package request

import (
	"vm-backend/pkg/helpers/db"
)

type ChangeRole struct {
	ID       uint  `json:"id" query:"id" validate:"required"`
	BranchID *uint `json:"branch_id,omitempty" query:"branch_id"`
	RoleID   uint  `json:"role_id" validate:"required"`
}

func (r *ChangeRole) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID).
		PtrWhere("branch_id = ?", r.BranchID)
}

func (r *ChangeRole) ToRoleQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.RoleID)
}

func (r *ChangeRole) ToUpdate() map[string]interface{} {
	return map[string]interface{}{
		"role_id": r.RoleID,
	}
}
