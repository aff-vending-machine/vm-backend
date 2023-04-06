package request

import (
	"encoding/json"
	"fmt"
)

type Update struct {
	MachineID uint    `json:"machine_id" query:"machine_id" validate:"required"`
	ID        uint    `json:"id,omitempty" query:"id" validate:"required"`
	ProductID *uint   `json:"product_id,omitempty"`
	Code      *string `json:"code,omitempty"`
	Stock     *uint   `json:"stock,omitempty"`
	Capacity  *uint   `json:"capacity,omitempty"`
	IsEnable  *bool   `json:"is_enable,omitempty"`
}

func (r *Update) ToMachineFilter() []string {
	return []string{
		fmt.Sprintf("id:=:%d", r.MachineID),
	}
}

func (r *Update) ToFilter() []string {
	return []string{
		fmt.Sprintf("machine_id:=:%d", r.MachineID),
		fmt.Sprintf("id:=:%d", r.ID),
	}
}

func (r *Update) ToJson() map[string]interface{} {
	var data map[string]interface{}

	b, _ := json.Marshal(r)
	json.Unmarshal(b, &data)

	delete(data, "machine_id")
	delete(data, "id")
	return data
}
