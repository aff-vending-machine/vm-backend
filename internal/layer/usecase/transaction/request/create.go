package request

import (
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Create struct {
	MerchantOrderID     string    `json:"merchant_order_id" validate:"required"`
	MachineSerialNumber string    `json:"machine_serial_number" validate:"required"`
	RawCart             string    `json:"raw_cart" validate:"required"`
	OrderQuantity       int       `json:"order_quantity" validate:"int|min:0"`
	OrderPrice          float64   `json:"order_price" validate:"int|min:0"`
	OrderedAt           time.Time `json:"ordered_at"`
	OrderStatus         string    `json:"order_status"`
	PaymentChannel      string    `json:"payment_channel,omitempty"`
	Reference1          string    `json:"reference1,omitempty"`
	Reference2          string    `json:"reference2,omitempty"`
	Reference3          string    `json:"reference3,omitempty"`
}

func (r *Create) ToEntity() *entity.Transaction {
	return &entity.Transaction{
		MerchantOrderID: r.MerchantOrderID,
	}
}
