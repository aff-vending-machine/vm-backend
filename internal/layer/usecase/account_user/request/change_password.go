package request

import (
	"vm-backend/pkg/helpers/db"
)

type ChangePassword struct {
	ID          uint   `json:"id" query:"id" validate:"required"`
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

func (r *ChangePassword) ToQuery() *db.Query {
	return db.NewQuery().AddWhere("id = ?", r.ID)
}

func (r *ChangePassword) ToUpdate(hashed string) map[string]interface{} {
	return map[string]interface{}{
		"password": hashed,
	}
}
