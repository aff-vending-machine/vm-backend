package request

import "vm-backend/pkg/helpers/db"

type Get struct {
	MachineID uint    `json:"machine_id" query:"machine_id" validate:"required"`
	ID        uint    `json:"id" query:"id" validate:"required"`
	Preloads  *string `json:"preloads,omitempty" query:"preloads"`
}

func (r *Get) ToQuery() *db.Query {
	return db.NewQuery().
		Where("machine_id = ?", r.MachineID).
		Where("id = ?", r.ID).
		PreloadsIf(r.Preloads)
}
