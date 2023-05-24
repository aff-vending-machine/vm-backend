package request

import "vm-backend/pkg/db"

type Delete struct {
	ID      uint  `json:"id" query:"id" validate:"required"`
	GroupID *uint `json:"group_id,omitempty" query:"group_id"`
}

func (r *Delete) ToQuery() *db.Query {
	return db.NewQuery().
		PtrWhere("group_id = ?", r.GroupID).
		AddWhere("id = ?", r.ID)
}
