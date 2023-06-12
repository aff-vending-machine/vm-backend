package request

import (
	"vm-backend/pkg/helpers/conv"
	"vm-backend/pkg/helpers/db"
)

type Update struct {
	MachineID uint  `json:"machine_id" query:"machine_id" validate:"required"`
	ID        uint  `json:"id,omitempty" query:"id" validate:"required"`
	ProductID *uint `json:"product_id,omitempty"`
	Stock     *uint `json:"stock,omitempty"`
	Capacity  *uint `json:"capacity,omitempty"`
	IsEnable  *bool `json:"is_enable,omitempty"`
}

func (r *Update) ToQuery() *db.Query {
	return db.NewQuery().
		Where("machine_id = ?", r.MachineID).
		Where("id = ?", r.ID)
}

func (r *Update) ToUpdate() map[string]interface{} {
	result, _ := conv.StructToMap(r)
	delete(result, "machine_id")
	delete(result, "id")
	return result
}
