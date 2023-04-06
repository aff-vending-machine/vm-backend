package response

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type PaymentChannel struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Channel      string `json:"channel"`
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

func PaymentChannelEntityToView(e *entity.PaymentChannel) *PaymentChannel {
	return &PaymentChannel{
		ID:           e.ID,
		Name:         e.Name,
		Channel:      e.Channel,
		Vendor:       e.Vendor,
		Active:       e.Active,
		Host:         e.Host,
		MerchantID:   e.MerchantID,
		MerchantName: e.MerchantName,
		BillerCode:   e.BillerCode,
		BillerID:     e.BillerID,
		StoreID:      e.StoreID,
		TerminalID:   e.TerminalID,
	}
}

func PaymentChannelEntityToList(es []entity.PaymentChannel) []PaymentChannel {
	items := make([]PaymentChannel, len(es))
	for i, e := range es {
		items[i] = *PaymentChannelEntityToView(&e)
	}

	return items
}
