package request

import (
	"time"
	"vm-backend/pkg/db"
)

type Report struct {
	MachineID uint       `json:"machine_id" query:"machine_id" validate:"required"`
	ChannelID *uint      `json:"channel_id,omitempty" query:"channel_id"`
	SortBy    *string    `json:"sort_by,omitempty" query:"sort_by"`
	From      *time.Time `json:"from,omitempty"`
	To        *time.Time `json:"to,omitempty"`
}

func (r *Report) ToTransactionQuery() *db.Query {
	query := db.NewQuery().
		AddWhere("order_status = ?", "DONE").
		AddWhere("machine_id = ?", r.MachineID).
		PtrWhere("channel_id = ?", r.ChannelID).
		PtrOrder(r.SortBy).
		AddPreload("Machine").
		AddPreload("Channel")
	if r.From != nil && r.To != nil {
		query = query.AddWhere("confirmed_paid_at BETWEEN ? AND ?", r.From, r.To)
	} else if r.From != nil {
		query = query.AddWhere("confirmed_paid_at >= ?", r.From)
	} else if r.To != nil {
		query = query.AddWhere("confirmed_paid_at <= ?", r.From)
	}

	return query
}
