package request

import (
	"vm-backend/pkg/helpers/db"
)

type GetAccountPermission struct {
	UserID uint   `json:"user_id" query:"id" validate:"required"`
	Scope  string `json:"scope" validate:"required"`
}

func (r *GetAccountPermission) ToUserQuery() *db.Query {
	return db.NewQuery().
		Where("id = ?", r.UserID).
		Preload("Role").
		Preload("Role.Permissions")
}
