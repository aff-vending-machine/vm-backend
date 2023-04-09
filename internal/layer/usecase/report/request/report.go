package request

import (
	"fmt"
	"time"
)

type Report struct {
	MachineID *uint      `json:"machine_id,omitempty"`
	From      *time.Time `json:"from,omitempty"`
	To        *time.Time `json:"to,omitempty"`
	Available *bool      `json:"available,omitempty"`
}

func (r *Report) ToSlotFilter() []string {
	filter := []string{
		"code||SORT||asc",
	}

	if r.MachineID != nil {
		filter = append(filter, fmt.Sprintf("machine_id||=||%d", *r.MachineID))
	}

	return filter
}

func (r *Report) ToPaymentFilter() []string {
	layout := "2006-01-02 15:04:05"

	filter := []string{
		"order_status||=||DONE",
		"confirmed_paid_at||SORT||asc",
	}

	if r.MachineID != nil {
		filter = append(filter, fmt.Sprintf("machine_id||=||%d", *r.MachineID))
	}

	if r.From != nil && r.To != nil {
		filter = append(filter, fmt.Sprintf("confirmed_paid_at||BETWEEN||%s,%s||time", r.From.Format(layout), r.To.Format(layout)))
	} else if r.From != nil {
		filter = append(filter, fmt.Sprintf("confirmed_paid_at||>||%s||time", r.From.Format(layout)))
	} else if r.To != nil {
		filter = append(filter, fmt.Sprintf("confirmed_paid_at||<||%s||time", r.To.Format(layout)))
	}

	return filter
}
