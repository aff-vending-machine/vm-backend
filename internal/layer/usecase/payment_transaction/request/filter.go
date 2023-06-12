package request

import (
	"time"
	"vm-backend/pkg/helpers/db"

	"github.com/rs/zerolog/log"
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
		LimitIfNotNil(r.Limit).
		OffsetIf(r.Offset).
		OrderIf(r.SortBy).
		WhereIf("id = ?", r.ID).
		WhereIf("branch_id = ?", r.BranchID).
		WhereIf("machine_id = ?", r.MachineID).
		WhereIf("channel_id = ?", r.ChannelID).
		WhereIf("merchant_order_id = ?", r.MerchantOrderID).
		WhereIf("order_status = ?", r.OrderStatus).
		PreloadsIf(r.Preloads)

	if r.From != nil && r.To != nil {
		query = query.Where("confirmed_paid_at BETWEEN ? AND ?", r.From, r.To)
	} else if r.From != nil {
		query = query.Where("confirmed_paid_at >= ?", r.From)
	} else if r.To != nil {
		query = query.Where("confirmed_paid_at <= ?", r.From)
	}

	log.Debug().Interface("query", query).Msg("query")

	return query
}
