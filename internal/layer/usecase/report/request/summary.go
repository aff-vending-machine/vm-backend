package request

import (
	"time"
	"vm-backend/pkg/helpers/db"
)

type Summary struct {
	BranchID *uint      `json:"branch_id,omitempty" query:"branch_id"`
	SortBy   *string    `json:"sort_by,omitempty" query:"sort_by"`
	From     *time.Time `json:"from,omitempty"`
	To       *time.Time `json:"to,omitempty"`
}

func (r *Summary) ToMachineQuery() *db.Query {
	return db.NewQuery().
		SetOrder("id:asc").
		PtrWhere("branch_id = ?", r.BranchID)
}

func (r *Summary) ToChannelQuery() *db.Query {
	return db.NewQuery().
		SetOrder("id:asc")
}

func (r *Summary) ToTransactionQuery() *db.Query {
	query := db.NewQuery().
		AddWhere("order_status = ?", "DONE").
		PtrWhere("branch_id = ?", r.BranchID).
		PtrOrder(r.SortBy)
	if r.From != nil && r.To != nil {
		query = query.AddWhere("confirmed_paid_at BETWEEN ? AND ?", r.From, r.To)
	} else if r.From != nil {
		query = query.AddWhere("confirmed_paid_at >= ?", r.From)
	} else if r.To != nil {
		query = query.AddWhere("confirmed_paid_at <= ?", r.From)
	}

	return query
}
