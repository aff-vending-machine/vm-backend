package sync

import (
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Transaction struct {
	ID                  uint       `json:"id"`
	MerchantOrderID     string     `json:"merchant_order_id"`     // key to find order
	MachineSerialNumber string     `json:"machine_serial_number"` // key to find machine
	Location            string     `json:"location"`              // ordered
	RawCart             string     `json:"raw_cart"`              // ordered
	OrderQuantity       int        `json:"order_quantity"`        // ordered
	OrderPrice          float64    `json:"order_price"`           // ordered
	OrderStatus         string     `json:"order_status"`          // ordered
	OrderedAt           time.Time  `json:"ordered_at"`            // ordered
	PaymentChannel      string     `json:"payment_channel"`       // ordered, key to find payment channel - MakeTransactionCreateRequest
	PaymentRequestedAt  *time.Time `json:"payment_requested_at"`  // ordered - MakeTransactionCreateRequest
	Reference1          *string    `json:"reference1"`            // reference1 - MakeTransactionCreateResult
	Reference2          *string    `json:"reference2"`            // reference2
	Reference3          *string    `json:"reference3"`            // reference3
	CancelledBy         *string    `json:"cancelled_by"`          // cancelled - MakeTransactionCancel
	CancelledAt         *time.Time `json:"cancelled_at"`          // cancelled - MakeTransactionCancel
	ConfirmedPaidBy     *string    `json:"confirmed_paid_by"`     // paid - MakeTransactionPaid
	ConfirmedPaidAt     *time.Time `json:"confirmed_paid_at"`     // paid - MakeTransactionPaid
	RefError            *string    `json:"ref_error"`             // MakeTransactionError
	RefundAt            *time.Time `json:"refund_at"`             // refund
	RefundPrice         float64    `json:"refund_price"`          // refund
	ReceivedItemAt      *time.Time `json:"received_item_at"`      // received - MakeTransactionDone
	ReceivedQuantity    int        `json:"received_quantity"`     // received, refund - MakeTransactionDone
	PaidPrice           float64    `json:"paid_price"`            // received, refund - MakeTransactionDone
	IsError             bool       `json:"is_error"`              // error
	Error               *string    `json:"error"`                 // error - MakeTransactionError
	ErrorAt             *time.Time `json:"error_at"`              // MakeTransactionRefund
}

func (m *Transaction) ToEntity(machineID uint, name string) *entity.Transaction {
	return &entity.Transaction{
		MachineID:           machineID,
		MachineName:         name,
		MerchantOrderID:     m.MerchantOrderID,
		MachineSerialNumber: m.MachineSerialNumber,
		Location:            m.Location,
		RawCart:             m.RawCart,
		OrderQuantity:       m.OrderQuantity,
		OrderPrice:          m.OrderPrice,
		OrderStatus:         m.OrderStatus,
		OrderedAt:           m.OrderedAt,
		PaymentChannel:      m.PaymentChannel,
		PaymentRequestedAt:  m.PaymentRequestedAt,
		Reference1:          m.Reference1,
		Reference2:          m.Reference2,
		Reference3:          m.Reference3,
		CancelledBy:         m.CancelledBy,
		CancelledAt:         m.CancelledAt,
		ConfirmedPaidBy:     m.ConfirmedPaidBy,
		ConfirmedPaidAt:     m.ConfirmedPaidAt,
		RefError:            m.RefError,
		RefundAt:            m.RefundAt,
		RefundPrice:         m.RefundPrice,
		ReceivedItemAt:      m.ReceivedItemAt,
		ReceivedQuantity:    m.ReceivedQuantity,
		PaidPrice:           m.PaidPrice,
		IsError:             m.IsError,
		Error:               m.Error,
		ErrorAt:             m.ErrorAt,
	}
}

func (m *Transaction) ToUpdate() map[string]interface{} {
	return map[string]interface{}{
		"order_quantity":       m.OrderQuantity,
		"order_price":          m.OrderPrice,
		"order_status":         m.OrderStatus,
		"ordered_at":           m.OrderedAt,
		"payment_channel":      m.PaymentChannel,
		"payment_requested_at": m.PaymentRequestedAt,
		"reference1":           m.Reference1,
		"reference2":           m.Reference2,
		"reference3":           m.Reference3,
		"cancelled_by":         m.CancelledBy,
		"cancelled_at":         m.CancelledAt,
		"confirmed_paid_by":    m.ConfirmedPaidBy,
		"confirmed_paid_at":    m.ConfirmedPaidAt,
		"ref_error":            m.RefError,
		"refund_at":            m.RefundAt,
		"refund_price":         m.RefundPrice,
		"received_item_at":     m.ReceivedItemAt,
		"received_quantity":    m.ReceivedQuantity,
		"paid_price":           m.PaidPrice,
		"is_error":             m.IsError,
		"error":                m.Error,
		"error_at":             m.ErrorAt,
	}
}
