package request

import "vm-backend/pkg/helpers/db"

type ValidateToken struct {
	Token string `json:"token" validate:"required"`
}

func (r *ValidateToken) ToUserQuery(id uint) *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", id).
		AddPreload("Role").
		AddPreload("Role.Permissions")
}
