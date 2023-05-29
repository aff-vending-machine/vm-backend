package models

import (
	"vm-backend/internal/core/domain/machine"
)

type Slot struct {
	Code     string   `json:"code"`
	Stock    int      `json:"stock"`
	Capacity int      `json:"capacity"`
	Product  *Product `json:"product,omitempty"`
	IsEnable bool     `json:"is_enable"`
}

func (m *Slot) ToDomain(machineID, productID uint) *machine.Slot {
	return &machine.Slot{
		MachineID: machineID,
		ProductID: productID,
		Code:      m.Code,
		Stock:     m.Stock,
		Capacity:  m.Capacity,
		IsEnable:  m.IsEnable,
	}
}

func (m *Slot) ToUpdate(productID uint) map[string]interface{} {
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
