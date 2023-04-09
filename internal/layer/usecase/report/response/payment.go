package response

import (
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Payment struct {
	ID                 uint       `json:"id"`
	MerchantOrderID    string     `json:"merchant_order_id"`
	MachineID          uint       `json:"machine_id"`
	PaymentChannel     string     `json:"payment_channel"`
	PaymentRequestedAt *time.Time `json:"payment_requested_at"`
	ConfirmedPaidBy    *string    `json:"confirmed_paid_by"`
	ConfirmedPaidAt    *time.Time `json:"confirmed_paid_at"`
	CancelledBy        *string    `json:"cancelled_by"`
	CancelledAt        *time.Time `json:"cancelled_at"`
	RefundAt           *time.Time `json:"refund_at"`
	RefundPrice        float64    `json:"refund_price"`
	PaidPrice          float64    `json:"paid_price"`
	Error              *string    `json:"error"`
	ErrorAt            *time.Time `json:"error_at"`
	OrderStatus        string     `json:"order_status"`
	Reference1         *string    `json:"reference1"`
	Reference2         *string    `json:"reference2"`
	Reference3         *string    `json:"reference3"`
}

func ToPayment(e *entity.Transaction) *Payment {
	return &Payment{
		ID:                 e.ID,
		MerchantOrderID:    e.MerchantOrderID,
		PaymentChannel:     e.PaymentChannel,
		PaymentRequestedAt: e.PaymentRequestedAt,
		ConfirmedPaidBy:    e.ConfirmedPaidBy,
		ConfirmedPaidAt:    e.ConfirmedPaidAt,
		CancelledBy:        e.CancelledBy,
		CancelledAt:        e.CancelledAt,
		RefundAt:           e.RefundAt,
		RefundPrice:        e.RefundPrice,
		PaidPrice:          e.PaidPrice,
		Error:              e.Error,
		ErrorAt:            e.ErrorAt,
		OrderStatus:        e.OrderStatus,
		Reference1:         e.Reference1,
		Reference2:         e.Reference2,
		Reference3:         e.Reference3,
	}
}

func ToPaymentList(ps []entity.Transaction) []Payment {
	items := make([]Payment, len(ps))
	for i, p := range ps {
		items[i] = *ToPayment(&p)
	}

	return items
}
