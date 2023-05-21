package request

import (
	"vm-backend/pkg/db"
)

type Delete struct {
	MachineID uint `json:"machine_id" query:"machine_id" validate:"required"`
	ID        uint `json:"id" query:"id" validate:"required"`
}

func (r *Delete) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("machine_id = ?", r.MachineID).
		AddWhere("id = ?", r.ID)
}
