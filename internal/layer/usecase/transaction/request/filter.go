package request

import (
	"fmt"
	"strings"
)

type Filter struct {
	Limit           *int    `json:"limit,omitempty" query:"limit"`
	Offset          *int    `json:"offset,omitempty" query:"offset"`
	SortBy          *string `json:"sort_by,omitempty" query:"sort_by"`
	ID              *uint   `json:"id,omitempty" query:"id"`
	MachineID       *uint   `json:"machine_id" query:"machine_id"`
	MerchantOrderID *string `json:"merchant_order_id,omitempty" query:"merchant_order_id"`
	Location        *string `json:"location,omitempty" query:"location"`
	OrderStatus     *string `json:"order_status,omitempty" query:"order_status"`
	PaymentChannel  *string `json:"payment_channel,omitempty" query:"payment_channel"`
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

	if r.MerchantOrderID != nil {
		filter = append(filter, fmt.Sprintf("merchant_order_id||=||%s", *r.MerchantOrderID))
	}

	if r.Location != nil {
		filter = append(filter, fmt.Sprintf("location||=||%s", *r.Location))
	}

	if r.OrderStatus != nil {
		filter = append(filter, fmt.Sprintf("order_status||=||%s", *r.OrderStatus))
	}

	if r.PaymentChannel != nil {
		filter = append(filter, fmt.Sprintf("payment_channel||=||%s", *r.PaymentChannel))
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
