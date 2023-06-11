package request

import (
	"vm-backend/pkg/helpers/db"
)

type Sync struct {
	MachineID uint `json:"machine_id" query:"machine_id" validate:"required"`
}

func (r *Sync) ToSlotQuery() *db.Query {
	return db.NewQuery().
		AddWhere("machine_id = ?", r.MachineID).
		AddPreload("Product")
}

func (r *Sync) ToChannelQuery() *db.Query {
	return db.NewQuery().
		AddWhere("machine_id = ?", r.MachineID)
}

func (r *Sync) ToMachineQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.MachineID).
		AddPreload("Branch").
		AddPreload("Slots").
		AddPreload("Slots.Product")
}
