package request

import (
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Create struct {
	MachineID uint   `json:"machine_id" query:"machine_id" validate:"required"`
	Code      string `json:"code" validate:"required"`
	Stock     int    `json:"stock,omitempty" validate:"min:0"`
	Capacity  int    `json:"capacity,omitempty" validate:"min:0"`
	IsEnable  bool   `json:"is_enable,omitempty"`
	ProductID uint   `json:"product_id,omitempty"`
}

func (r *Create) ToMachineFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.MachineID),
	}
}

func (r *Create) ToEntity() *entity.MachineSlot {
	return &entity.MachineSlot{
		Code:      r.Code,
		Stock:     r.Stock,
		Capacity:  r.Capacity,
		IsEnable:  r.IsEnable,
		ProductID: r.ProductID,
	}
}
