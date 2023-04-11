package request

import (
	"fmt"
	"strings"
)

type Filter struct {
	Limit         *int    `json:"limit,omitempty" query:"limit"`
	Offset        *int    `json:"offset,omitempty" query:"offset"`
	SortBy        *string `json:"sort_by,omitempty" query:"sort_by"`
	ID            *uint   `json:"id,omitempty" query:"id"`
	MachineID     *uint   `json:"machine_id" query:"machine_id"`
	OrderID       *string `json:"order_id,omitempty" query:"order_id"`
	TransactionID *uint   `json:"transaction_id,omitempty" query:"transaction_id"`
}

func (r *Filter) ToFilter() []string {
	filter := []string{}

	if r.Limit != nil {
		filter = append(filter, fmt.Sprintf("||LIMIT||%d", *r.Limit))
	}

	if r.Offset != nil {
		filter = append(filter, fmt.Sprintf("||OFFSET||%d", *r.Offset))
	}

	if r.ID != nil {
		filter = append(filter, fmt.Sprintf("id||=||%d", *r.ID))
	}

	if r.MachineID != nil {
		filter = append(filter, fmt.Sprintf("machine_id||=||%d", *r.MachineID))
	}

	if r.OrderID != nil {
		filter = append(filter, fmt.Sprintf("order_id||=||%s", *r.OrderID))
	}

	if r.TransactionID != nil {
		filter = append(filter, fmt.Sprintf("transaction_id||=||%d", *r.TransactionID))
	}

	if r.SortBy != nil {
		val := strings.Split(*r.SortBy, ":")
		if len(val) == 1 || (val[1] != "asc" && val[1] != "desc") {
			filter = append(filter, fmt.Sprintf("%s||SORT||%s", val[0], "asc"))
		} else {
			filter = append(filter, fmt.Sprintf("%s||SORT||%s", val[0], val[1]))
		}
	} else {
		filter = append(filter, "ordered_at||SORT||desc")
	}

	return filter
}