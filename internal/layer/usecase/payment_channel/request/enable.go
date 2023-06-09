package request

import "vm-backend/pkg/helpers/db"

type Enable struct {
	ID uint `json:"id" query:"id" validate:"required"`
}

func (r *Enable) ToQuery() *db.Query {
	return db.NewQuery().
		Where("id = ?", r.ID)
}

func (r *Enable) ToUpdate() map[string]interface{} {
	return map[string]interface{}{
		"is_enable": true,
	}
}
