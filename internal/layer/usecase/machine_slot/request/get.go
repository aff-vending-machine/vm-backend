package request

import "vm-backend/pkg/db"

type Get struct {
	MachineID uint    `json:"machine_id" query:"machine_id" validate:"required"`
	ID        uint    `json:"id" query:"id" validate:"required"`
	Preloads  *string `json:"preloads,omitempty" query:"preloads"`
}

func (r *Get) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("machine_id = ?", r.MachineID).
		AddWhere("id = ?", r.ID).
		PtrPreloads(r.Preloads)
}
