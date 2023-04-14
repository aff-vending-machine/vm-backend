package response

import (
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Transaction struct {
	ID                  uint       `json:"id"`
	MerchantOrderID     string     `json:"merchant_order_id"`
	MachineID           uint       `json:"machine_id"`
	MachineName         string     `json:"machine_name"`
	MachineSerialNumber string     `json:"machine_serial_number"`
	Location            string     `json:"location"`
	Note                string     `json:"note"`
	OrderQuantity       int        `json:"order_quantity"`
	OrderPrice          float64    `json:"order_price"`
	OrderStatus         string     `json:"order_status"`
	OrderedAt           time.Time  `json:"ordered_at"`
	PaymentChannel      string     `json:"payment_channel"`
	PaymentRequestedAt  *time.Time `json:"payment_requested_at"`
	Reference1          *string    `json:"reference1"`
	Reference2          *string    `json:"reference2"`
	Reference3          *string    `json:"reference3"`
	CancelledBy         *string    `json:"cancelled_by"`
	CancelledAt         *time.Time `json:"cancelled_at"`
	ConfirmedPaidBy     *string    `json:"confirmed_paid_by"`
	ConfirmedPaidAt     *time.Time `json:"confirmed_paid_at"`
	RefundAt            *time.Time `json:"refund_at"`
	RefundPrice         float64    `json:"refund_price"`
	ReceivedItemAt      *time.Time `json:"received_item_at"`
	ReceivedQuantity    int        `json:"received_quantity"`
	PaidPrice           float64    `json:"paid_price"`
	IsError             bool       `json:"is_error"`
	Error               *string    `json:"error"`
	ErrorAt             *time.Time `json:"error_at"`
}

func ToTransaction(e *entity.Transaction) *Transaction {
	return &Transaction{
		ID:                  e.ID,
		MerchantOrderID:     e.MerchantOrderID,
		MachineID:           e.MachineID,
		MachineName:         e.MachineName,
		MachineSerialNumber: e.MachineSerialNumber,
		Location:            e.Location,
		Note:                e.Note,
		OrderQuantity:       e.OrderQuantity,
		OrderPrice:          e.OrderPrice,
		OrderStatus:         e.OrderStatus,
		OrderedAt:           e.OrderedAt,
		PaymentChannel:      e.PaymentChannel,
		PaymentRequestedAt:  e.PaymentRequestedAt,
		Reference1:          e.Reference1,
		Reference2:          e.Reference2,
		Reference3:          e.Reference3,
		CancelledBy:         e.CancelledBy,
		CancelledAt:         e.CancelledAt,
		ConfirmedPaidBy:     e.ConfirmedPaidBy,
		ConfirmedPaidAt:     e.ConfirmedPaidAt,
		RefundAt:            e.RefundAt,
		RefundPrice:         e.RefundPrice,
		ReceivedItemAt:      e.ReceivedItemAt,
		ReceivedQuantity:    e.ReceivedQuantity,
		PaidPrice:           e.PaidPrice,
		IsError:             e.IsError,
		Error:               e.Error,
		ErrorAt:             e.ErrorAt,
	}
}

func ToTransactionList(ps []entity.Transaction) []Transaction {
	items := make([]Transaction, len(ps))
	for i, p := range ps {
		items[i] = *ToTransaction(&p)
	}

	return items
}
