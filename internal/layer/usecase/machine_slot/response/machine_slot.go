package response

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/product/response"
)

type MachineSlot struct {
	ID       uint              `json:"id"`
	Code     string            `json:"code"`
	Stock    int               `json:"stock"`
	Capacity int               `json:"capacity"`
	Product  *response.Product `json:"product"`
	IsEnable bool              `json:"is_enable"`
}

func ToMachineSlot(e *entity.MachineSlot) *MachineSlot {
	return &MachineSlot{
		ID:       e.ID,
		Code:     e.Code,
		Stock:    e.Stock,
		Capacity: e.Capacity,
		Product:  response.ToProduct(e.Product),
		IsEnable: e.IsEnable,
	}
}

func ToMachineSlotList(es []entity.MachineSlot) []MachineSlot {
	items := make([]MachineSlot, len(es))
	for i, e := range es {
		items[i] = *ToMachineSlot(&e)
	}

	return items
}
