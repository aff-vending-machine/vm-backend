package request

import "vm-backend/pkg/db"

type Delete struct {
	GroupID uint `json:"group_id"`
	ID      uint `json:"id" query:"id" validate:"required"`
}

func (r *Delete) ToQuery() *db.Query {
	return db.NewQuery().
		PtrWhere("group_id = ?", r.GroupID).
		AddWhere("id = ?", r.ID)
}
