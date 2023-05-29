package models

import (
	"vm-backend/internal/core/domain/catalog"
	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/core/domain/payment"
)

func FromMachine(e *machine.Machine) *Machine {
	branch := "-"
	if e.Branch != nil {
		branch = e.Branch.Name
	}

	return &Machine{
		Name:         e.Name,
		SerialNumber: e.SerialNumber,
		Branch:       branch,
		Location:     e.Location,
		Type:         e.Type,
		Vendor:       e.Vendor,
		Status:       e.Status,
	}
}

func FromSlot(e *machine.Slot) *Slot {
	return &Slot{
		Code:     e.Code,
		Stock:    e.Stock,
		Capacity: e.Capacity,
		Product:  FromProduct(e.Product),
		IsEnable: e.IsEnable,
	}
}

func FromSlotList(entities []machine.Slot) []Slot {
	results := make([]Slot, len(entities))
	for i, e := range entities {
		results[i] = *FromSlot(&e)
	}

	return results
}

func FromProduct(e *catalog.Product) *Product {
	if e == nil {
		return nil
	}

	return &Product{
		SKU:      e.SKU,
		Name:     e.Name,
		ImageURL: e.ImageURL,
		Price:    e.SalePrice,
	}
}

func FromChannel(e *payment.Channel) *Channel {
	return &Channel{
		Channel:      e.Channel,
		Name:         e.Name,
		Vendor:       e.Vendor,
		IsEnable:     e.IsEnable,
		Host:         e.Host,
		MerchantID:   e.MerchantID,
		MerchantName: e.MerchantName,
		BillerCode:   e.BillerCode,
		BillerID:     e.BillerID,
		Token:        e.Token,
		StoreID:      e.StoreID,
		TerminalID:   e.TerminalID,
	}
}

func FromChannelList(entities []payment.Channel) []Channel {
	results := make([]Channel, len(entities))
	for i, e := range entities {
		results[i] = *FromChannel(&e)
	}

	return results
}
