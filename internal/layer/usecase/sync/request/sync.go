package request

import (
	"fmt"
)

type Sync struct {
	MachineID uint `json:"machine_id" query:"machine_id" validate:"required"`
}

func (r *Sync) ToSlotFilter() []string {
	return []string{
		fmt.Sprintf("machine_id||=||%d", r.MachineID),
		"||PRELOAD||Product",
	}
}

func (r *Sync) ToMachineFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.MachineID),
		"||PRELOAD||Slots",
	}
}
