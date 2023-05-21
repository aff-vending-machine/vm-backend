package request

import "vm-backend/pkg/db"

type Disable struct {
	ID uint `json:"id" query:"id" validate:"required"`
}

func (r *Disable) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID)
}

func (r *Disable) ToUpdate() map[string]interface{} {
	return map[string]interface{}{
		"is_enable": false,
	}
}
