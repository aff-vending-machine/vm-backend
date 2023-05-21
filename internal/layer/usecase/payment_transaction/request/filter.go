package request

import (
	"time"
	"vm-backend/pkg/db"
)

type Filter struct {
	Limit           *int       `json:"limit,omitempty" query:"limit"`
	Offset          *int       `json:"offset,omitempty" query:"offset"`
	SortBy          *string    `json:"sort_by,omitempty" query:"sort_by"`
	Preloads        *string    `json:"preloads,omitempty" query:"preloads"`
	ID              *uint      `json:"id,omitempty" query:"id"`
	BranchID        *uint      `json:"branch_id,omitempty" query:"branch_id"`
	MachineID       *uint      `json:"machine_id,omitempty" query:"machine_id"`
	ChannelID       *uint      `json:"channel_id,omitempty" query:"channel_id"`
	MerchantOrderID *string    `json:"merchant_order_id,omitempty" query:"merchant_order_id"`
	OrderStatus     *string    `json:"order_status,omitempty" query:"order_status"`
	From            *time.Time `json:"from,omitempty"`
	To              *time.Time `json:"to,omitempty"`
}

func (r *Filter) ToQuery() *db.Query {
	query := db.NewQuery().
		PtrLimit(r.Limit).
		PtrOffset(r.Offset).
		PtrOrder(r.SortBy).
		PtrWhere("id = ?", r.ID).
		PtrWhere("branch_id = ?", r.BranchID).
		PtrWhere("machine_id = ?", r.MachineID).
		PtrWhere("channel_id = ?", r.ChannelID).
		PtrWhere("merchant_order_id = ?", r.MerchantOrderID).
		PtrWhere("order_status = ?", r.OrderStatus).
		PtrPreloads(r.Preloads)

	if r.From != nil && r.To != nil {
		query = query.AddWhere("confirmed_paid_at BETWEEN ? AND ?", r.From, r.To)
	} else if r.From != nil {
		query = query.AddWhere("confirmed_paid_at >= ?", r.From)
	} else if r.To != nil {
		query = query.AddWhere("confirmed_paid_at <= ?", r.From)
	}

	return query
}
