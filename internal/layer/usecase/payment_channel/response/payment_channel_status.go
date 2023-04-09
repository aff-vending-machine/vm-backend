package response

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type PaymentChannelStatus struct {
	ID     uint `json:"id"`
	Active bool `json:"active"`
}

func ToPaymentChannelStatus(e *entity.PaymentChannel) *PaymentChannelStatus {
	return &PaymentChannelStatus{
		ID:     e.ID,
		Active: e.Active,
	}
}
