package model

import "github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"

type Slot struct {
	Code     string   `json:"code"`
	Stock    int      `json:"stock"`
	Capacity int      `json:"capacity"`
	Product  *Product `json:"product,omitempty"`
	IsEnable bool     `json:"is_enable"`
}

func ToSlot(e *entity.MachineSlot) *Slot {
	return &Slot{
		Code:     e.Code,
		Stock:    e.Stock,
		Capacity: e.Capacity,
		Product:  ToProduct(e.Product),
		IsEnable: e.IsEnable,
	}
}

func ToSlotList(entities []entity.MachineSlot) []Slot {
	results := make([]Slot, len(entities))
	for i, e := range entities {
		results[i] = *ToSlot(&e)
	}

	return results
}

func (m *Slot) ToEntity(machineID, productID uint) *entity.MachineSlot {
	return &entity.MachineSlot{
		MachineID: machineID,
		ProductID: productID,
		Code:      m.Code,
		Stock:     m.Stock,
		Capacity:  m.Capacity,
		IsEnable:  m.IsEnable,
	}
}

func (m *Slot) ToJson(productID uint) map[string]interface{} {
	slot := map[string]interface{}{
		"code":      m.Code,
		"stock":     m.Stock,
		"capacity":  m.Capacity,
		"is_enable": m.IsEnable,
	}

	if m.Product != nil {
		slot["product_id"] = productID
	}

	return slot
}
