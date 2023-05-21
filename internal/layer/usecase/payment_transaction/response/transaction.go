package response

import (
	"time"
)

type Transaction struct {
	ID                 uint       `json:"id"`
	BranchID           uint       `json:"branch_id"`
	Branch             Branch     `json:"branch"`
	MachineID          uint       `json:"machine_id"`
	Machine            Machine    `json:"machine"`
	ChannelID          uint       `json:"channel_id"`
	Channel            Channel    `json:"channel"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	MerchantOrderID    string     `json:"merchant_order_id"`
	RawCart            string     `json:"raw_cart"`
	Note               string     `json:"note"`
	OrderQuantity      int        `json:"order_quantity"`
	OrderPrice         float64    `json:"order_price"`
	OrderStatus        string     `json:"order_status"`
	OrderedAt          time.Time  `json:"ordered_at"`
	PaymentRequestedAt *time.Time `json:"payment_requested_at"`
	Reference1         *string    `json:"reference1"`
	Reference2         *string    `json:"reference2"`
	Reference3         *string    `json:"reference3"`
	CancelledBy        *string    `json:"cancelled_by"`
	CancelledAt        *time.Time `json:"cancelled_at"`
	ConfirmedPaidBy    *string    `json:"confirmed_paid_by"`
	ConfirmedPaidAt    *time.Time `json:"confirmed_paid_at"`
	RefundAt           *time.Time `json:"refund_at"`
	RefundPrice        float64    `json:"refund_price"`
	ReceivedItemAt     *time.Time `json:"received_item_at"`
	ReceivedQuantity   int        `json:"received_quantity"`
	PaidPrice          float64    `json:"paid_price"`
	IsError            bool       `json:"is_error"`
	Error              *string    `json:"error"`
	ErrorAt            *time.Time `json:"error_at"`
}
