package response

import (
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Payment struct {
	MerchantOrderID     string     `json:"merchant_order_id"`
	MachineID           uint       `json:"machine_id"`
	MachineName         string     `json:"machine_name"`
	MachineSerialNumber string     `json:"machine_serial_number"`
	Location            string     `json:"location"`
	PaymentChannel      string     `json:"payment_channel"`
	ConfirmedPaidBy     *string    `json:"confirmed_paid_by"`
	OrderedAt           time.Time  `json:"ordered_at"`
	PaymentRequestedAt  *time.Time `json:"payment_requested_at"`
	ConfirmedPaidAt     *time.Time `json:"paid_at"`
	ReceivedItemAt      *time.Time `json:"received_item_at" `
	OrderQuantity       int        `json:"order_quantity"`
	ReceivedQuantity    int        `json:"received_quantity"`
	OrderPrice          float64    `json:"order_price"`
	PaidPrice           float64    `json:"paid_price"`
	Reference1          *string    `json:"reference1"`
	Reference2          *string    `json:"reference2"`
	Reference3          *string    `json:"reference3"`
	Note                string     `json:"note"`
}

func ToPayment(e *entity.Transaction) *Payment {
	return &Payment{
		MerchantOrderID:     e.MerchantOrderID,
		MachineID:           e.MachineID,
		MachineName:         e.MachineName,
		MachineSerialNumber: e.MachineSerialNumber,
		Location:            e.Location,
		PaymentChannel:      e.PaymentChannel,
		ConfirmedPaidBy:     e.ConfirmedPaidBy,
		OrderedAt:           e.OrderedAt,
		PaymentRequestedAt:  e.PaymentRequestedAt,
		ConfirmedPaidAt:     e.ConfirmedPaidAt,
		ReceivedItemAt:      e.ReceivedItemAt,
		OrderQuantity:       e.OrderQuantity,
		ReceivedQuantity:    e.ReceivedQuantity,
		OrderPrice:          e.OrderPrice,
		PaidPrice:           e.PaidPrice,
		Reference1:          e.Reference1,
		Reference2:          e.Reference2,
		Reference3:          e.Reference3,
		Note:                e.Note,
	}
}

func ToPaymentList(ps []entity.Transaction) []Payment {
	items := make([]Payment, len(ps))
	for i, p := range ps {
		items[i] = *ToPayment(&p)
	}

	return items
}
