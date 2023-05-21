package request

import "vm-backend/pkg/db"

type Enable struct {
	ID uint `json:"id" query:"id" validate:"required"`
}

func (r *Enable) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID)
}

func (r *Enable) ToUpdate() map[string]interface{} {
	return map[string]interface{}{
		"is_enable": true,
	}
}
