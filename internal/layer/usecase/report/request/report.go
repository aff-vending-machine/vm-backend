package request

import (
	"time"
	"vm-backend/pkg/helpers/db"
)

type Report struct {
	MachineID uint       `json:"machine_id" query:"machine_id" validate:"required"`
	BranchID  *uint      `json:"branch_id,omitempty" query:"branch_id"`
	ChannelID *uint      `json:"channel_id,omitempty" query:"channel_id"`
	SortBy    *string    `json:"sort_by,omitempty" query:"sort_by"`
	From      *time.Time `json:"from,omitempty"`
	To        *time.Time `json:"to,omitempty"`
}

func (r *Report) ToTransactionQuery() *db.Query {
	query := db.NewQuery().
		Where("order_status = ?", "DONE").
		Where("machine_id = ?", r.MachineID).
		WhereIf("branch_id = ?", r.BranchID).
		WhereIf("channel_id = ?", r.ChannelID).
		OrderIf(r.SortBy).
		Preload("Machine").
		Preload("Channel")
	if r.From != nil && r.To != nil {
		query = query.Where("confirmed_paid_at BETWEEN ? AND ?", r.From, r.To)
	} else if r.From != nil {
		query = query.Where("confirmed_paid_at >= ?", r.From)
	} else if r.To != nil {
		query = query.Where("confirmed_paid_at <= ?", r.From)
	}

	return query
}
