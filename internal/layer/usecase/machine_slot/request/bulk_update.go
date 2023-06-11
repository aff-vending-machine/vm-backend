package request

import (
	"vm-backend/pkg/helpers/conv"
	"vm-backend/pkg/helpers/db"
)

type BulkUpdate struct {
	MachineID uint `json:"machine_id" query:"machine_id" validate:"required"`
	Data      []struct {
		ID        uint  `json:"id" validate:"required"`
		ProductID *uint `json:"product_id,omitempty"`
		Stock     *uint `json:"stock,omitempty"`
		Capacity  *uint `json:"capacity,omitempty"`
		IsEnable  *bool `json:"is_enable,omitempty"`
	} `json:"data" validate:"required"`
}

func (r *BulkUpdate) ToQuery(id uint) *db.Query {
	return db.NewQuery().
		AddWhere("machine_id = ?", r.MachineID).
		AddWhere("id = ?", id)
}

func (r *BulkUpdate) ToUpdate(index int) map[string]interface{} {
	result, _ := conv.StructToMap(r.Data[index])
	delete(result, "id")
	return result
}
