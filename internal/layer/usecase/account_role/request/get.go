package request

import "vm-backend/pkg/helpers/db"

type Get struct {
	ID       uint    `json:"id" query:"id" validate:"required"`
	Preloads *string `json:"preloads,omitempty" query:"preloads"`
}

func (r *Get) ToQuery() *db.Query {
	return db.NewQuery().
		Where("id = ?", r.ID).
		PreloadsIf(r.Preloads)
}
