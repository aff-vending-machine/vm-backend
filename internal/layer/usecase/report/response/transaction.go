package response

import (
	"time"
)

type Transaction struct {
	MerchantOrderID     string    `json:"merchant_order_id"`
	MachineID           uint      `json:"machine_id"`
	MachineName         string    `json:"machine_name"`
	MachineSerialNumber string    `json:"machine_serial_number"`
	Location            string    `json:"location"`
	PaymentChannel      string    `json:"payment_channel"`
	Cart                []Item    `json:"cart"`
	ConfirmedPaidBy     string    `json:"confirmed_paid_by"`
	OrderedAt           time.Time `json:"ordered_at"`
	PaymentRequestedAt  time.Time `json:"payment_requested_at"`
	ConfirmedPaidAt     time.Time `json:"confirmed_paid_at"`
	ReceivedItemAt      time.Time `json:"received_item_at" `
	OrderQuantity       int       `json:"order_quantity"`
	ReceivedQuantity    int       `json:"received_quantity"`
	OrderPrice          float64   `json:"order_price"`
	PaidPrice           float64   `json:"paid_price"`
	Reference1          string    `json:"reference1"`
	Reference2          string    `json:"reference2"`
	Reference3          string    `json:"reference3"`
	Note                string    `json:"note"`
}
