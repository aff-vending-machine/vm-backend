package request

import (
	"fmt"
)

type Get struct {
	MachineID uint `json:"machine_id" query:"machine_id" validate:"required"`
	ID        uint `json:"id" query:"id" validate:"required"`
}

func (r *Get) ToFilter() []string {
	return []string{
		fmt.Sprintf("machine_id:=:%d", r.MachineID),
		fmt.Sprintf("id:=:%d", r.ID),
		":PRELOAD:Product",
	}
}
