package request

import (
	"vm-backend/pkg/db"
)

type GetPermissionLevel struct {
	UserID uint   `json:"user_id" query:"id" validate:"required"`
	Scope  string `json:"scope" validate:"required"`
}

func (r *GetPermissionLevel) ToUserQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.UserID).
		AddPreload("Role").
		AddPreload("Role.Permissions")
}
