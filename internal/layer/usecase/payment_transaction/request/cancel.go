package request

import (
	"fmt"
	"time"
	"vm-backend/pkg/helpers/db"
)

type Cancel struct {
	ID       uint   `json:"id" query:"id" validate:"required"`
	Caller   string `json:"caller" query:"caller" validate:"required"`
	BranchID *uint  `json:"branch_id" query:"branch_id"`
}

func (r *Cancel) ToQuery() *db.Query {
	return db.NewQuery().
		Where("id = ?", r.ID).
		WhereIf("branch_id = ?", r.BranchID)
}

func (r *Cancel) ToUpdate() map[string]interface{} {
	return map[string]interface{}{
		"order_status":      "CANCELLED",
		"cancelled_by":      r.Caller,
		"cancelled_at":      time.Now(),
		"refund_at":         nil,
		"refund_price":      0,
		"received_item_at":  nil,
		"received_quantity": 0,
		"paid_price":        0,
		"note":              fmt.Sprintf("cancel order by %s", r.Caller),
		"confirmed_paid_by": nil,
		"confirmed_paid_at": nil,
	}
}
