package request

import (
	"vm-backend/pkg/helpers/db"
)

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *Login) ToUserQuery() *db.Query {
	return db.NewQuery().
		AddWhere("username = ?", r.Username).
		AddPreload("Role")
}
