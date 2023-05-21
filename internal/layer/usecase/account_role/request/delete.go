package request

import "vm-backend/pkg/db"

type Delete struct {
	ID uint `json:"id" query:"id" validate:"required"`
}

func (r *Delete) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID)
}
