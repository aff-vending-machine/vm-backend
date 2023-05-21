package request

import "vm-backend/pkg/db"

type ResetPassword struct {
	ID    uint `json:"id" query:"id" validate:"required"`
	Level int  `json:"level" query:"level" validate:"required"`
}

func (r *ResetPassword) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID)
}

func (r *ResetPassword) ToUpdate(hashed string) map[string]interface{} {
	return map[string]interface{}{
		"password": hashed,
	}
}
