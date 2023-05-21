package request

import "vm-backend/pkg/db"

type Get struct {
	ID         uint    `json:"id" query:"id" validate:"required"`
	Preloads   *string `json:"preloads,omitempty" query:"preloads"`
}

func (r *Get) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID).
		PtrPreloads(r.Preloads)
}
