package request

import (
	"fmt"
	"time"
	"vm-backend/pkg/helpers/db"
)

type Done struct {
	ID     uint   `json:"id" query:"id" validate:"required"`
	Caller string `json:"caller" query:"caller" validate:"required"`
}

func (r *Done) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID)
}

func (r *Done) ToUpdate(quantity int, price float64, confirmed *string) map[string]interface{} {
	now := time.Now()

	result := map[string]interface{}{
		"cancelled_by":      nil,
		"cancelled_at":      nil,
		"refund_at":         nil,
		"order_status":      "DONE",
		"refund_price":      0,
		"received_item_at":  now,
		"received_quantity": quantity,
		"paid_price":        price,
		"note":              fmt.Sprintf("confirm order by %s", r.Caller),
	}

	if confirmed == nil {
		result["confirmed_paid_by"] = fmt.Sprintf("user (%s)", r.Caller)
		result["confirmed_paid_at"] = now
	}

	return result
}
