package request

import (
	"encoding/json"
	"fmt"
)

type BulkUpdate struct {
	MachineID uint `json:"machine_id" query:"machine_id" validate:"required"`
	Data      []struct {
		ID        uint    `json:"id" validate:"required"`
		ProductID *uint   `json:"product_id,omitempty"`
		Stock     *uint   `json:"stock,omitempty"`
		Capacity  *uint   `json:"capacity,omitempty"`
		IsEnable  *bool   `json:"is_enable,omitempty"`
	}
}

func (r *BulkUpdate) ToMachineFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.MachineID),
	}
}

func (r *BulkUpdate) ToFilter(id uint) []string {
	return []string{
		fmt.Sprintf("machine_id||=||%d", r.MachineID),
		fmt.Sprintf("id||=||%d", id),
	}
}

func (r *BulkUpdate) ToJson(index int) map[string]interface{} {
	var data map[string]interface{}

	b, _ := json.Marshal(r.Data[index])
	json.Unmarshal(b, &data)

	delete(data, "id")
	return data
}
