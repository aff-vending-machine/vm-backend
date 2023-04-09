package request

import (
	"encoding/json"
	"fmt"
)

type Update struct {
	ID                  uint     `json:"id" query:"id" validate:"required"`
	MerchantOrderID     *string  `json:"merchant_order_id,omitempty"`
	MachineID           *uint    `json:"machine_id,omitempty"`
	MachineSerialNumber *string  `json:"machine_serial_number,omitempty"`
	OrderID             *string  `json:"order_id,omitempty"`
	OrderQuantity       *int     `json:"order_quantity,omitempty"`
	OrderPrice          *float64 `json:"order_price,omitempty"`
}

func (r *Update) ToFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.ID),
	}
}

func (r *Update) ToJson() map[string]interface{} {
	var data map[string]interface{}

	b, _ := json.Marshal(r)
	json.Unmarshal(b, &data)

	delete(data, "id")
	return data
}
