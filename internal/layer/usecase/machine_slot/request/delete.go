package request

import (
	"vm-backend/pkg/helpers/db"
)

type Delete struct {
	MachineID uint `json:"machine_id" query:"machine_id" validate:"required"`
	ID        uint `json:"id" query:"id" validate:"required"`
}

func (r *Delete) ToQuery() *db.Query {
	return db.NewQuery().
		Where("machine_id = ?", r.MachineID).
		Where("id = ?", r.ID)
}
