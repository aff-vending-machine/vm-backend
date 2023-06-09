package request

import "vm-backend/pkg/helpers/db"

type Delete struct {
	ID uint `json:"id" query:"id" validate:"required"`
}

func (r *Delete) ToQuery() *db.Query {
	return db.NewQuery().
		Where("id = ?", r.ID)
}
