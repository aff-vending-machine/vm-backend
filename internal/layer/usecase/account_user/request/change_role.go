package request

import (
	"vm-backend/pkg/db"
)

type ChangeRole struct {
	ID     uint `json:"id" query:"id" validate:"required"`
	RoleID uint `json:"role_id" validate:"required"`
}

func (r *ChangeRole) ToQuery() *db.Query {
	return db.NewQuery().AddWhere("id = ?", r.ID)
}

func (r *ChangeRole) ToRoleQuery() *db.Query {
	return db.NewQuery().AddWhere("id = ?", r.RoleID)
}

func (r *ChangeRole) ToUpdate() map[string]interface{} {
	return map[string]interface{}{
		"role_id": r.RoleID,
	}
}
