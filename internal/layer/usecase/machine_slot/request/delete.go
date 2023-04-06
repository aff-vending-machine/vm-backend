package request

import (
	"fmt"
)

type Delete struct {
	ID        uint `json:"id" query:"id" validate:"required"`
	MachineID uint `json:"machine_id" query:"machine_id" validate:"required"`
}

func (r *Delete) ToFilter() []string {
	return []string{
		fmt.Sprintf("id:=:%d", r.ID),
		fmt.Sprintf("machine_id:=:%d", r.MachineID),
	}
}

func (r *Delete) ToMachineFilter() []string {
	return []string{
		fmt.Sprintf("id:=:%d", r.MachineID),
	}
}
