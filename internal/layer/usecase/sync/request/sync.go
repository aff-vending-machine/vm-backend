package request

import (
	"vm-backend/pkg/helpers/db"
)

type Sync struct {
	MachineID uint `json:"machine_id" query:"machine_id" validate:"required"`
}

func (r *Sync) ToSlotQuery() *db.Query {
	return db.NewQuery().
		Where("machine_id = ?", r.MachineID).
		Preload("Product")
}

func (r *Sync) ToChannelQuery() *db.Query {
	return db.NewQuery().
		Where("machine_id = ?", r.MachineID)
}

func (r *Sync) ToMachineQuery() *db.Query {
	return db.NewQuery().
		Where("id = ?", r.MachineID).
		Preload("Branch").
		Preload("Slots").
		Preload("Slots.Product")
}
