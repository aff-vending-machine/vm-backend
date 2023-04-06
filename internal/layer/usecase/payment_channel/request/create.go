package request

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Create struct {
	Name         string `json:"name" validate:"required"`
	Channel      string `json:"channel" validate:"required"`
	Vendor       string `json:"vendor"`
	Active       bool   `json:"active"`
	Host         string `json:"host"`
	MerchantID   string `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	BillerCode   string `json:"biller_code"`
	BillerID     string `json:"biller_id"`
	StoreID      string `json:"store_id"`
	TerminalID   string `json:"terminal_id"`
}

func (r *Create) ToEntity() *entity.PaymentChannel {
	return &entity.PaymentChannel{
		Name:         r.Name,
		Channel:      r.Channel,
		Vendor:       r.Vendor,
		Active:       r.Active,
		Host:         r.Host,
		MerchantID:   r.MerchantID,
		MerchantName: r.MerchantName,
		BillerCode:   r.BillerCode,
		BillerID:     r.BillerID,
		StoreID:      r.StoreID,
		TerminalID:   r.TerminalID,
	}
}
