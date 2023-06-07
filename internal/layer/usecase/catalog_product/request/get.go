package request

import "vm-backend/pkg/helpers/db"

type Get struct {
	GroupID  uint    `json:"group_id" query:"group_id"`
	ID       uint    `json:"id" query:"id" validate:"required"`
	Preloads *string `json:"preloads,omitempty" query:"preloads"`
}

func (r *Get) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID).
		PtrWhere("group_id = ?", r.GroupID).
		PtrPreloads(r.Preloads)
}
