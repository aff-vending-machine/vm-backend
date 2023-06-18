package request

import "vm-backend/pkg/helpers/db"

type ValidateToken struct {
	Token string `json:"token" validate:"required"`
}

func (r *ValidateToken) ToUserQuery(id uint) *db.Query {
	return db.NewQuery().
		Where("id = ?", id).
		Preload("Role").
		Preload("Role.Permissions").
		Preload("Branch")
}
