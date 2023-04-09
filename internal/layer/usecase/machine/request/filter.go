package request

import (
	"fmt"
	"strings"
)

type Filter struct {
	Limit        *int    `json:"limit,omitempty" query:"limit"`
	Offset       *int    `json:"offset,omitempty" query:"offset"`
	SortBy       *string `json:"sort_by,omitempty" query:"sort_by"`
	ID           *uint   `json:"id,omitempty" query:"id"`
	Name         *string `json:"name,omitempty" query:"name"`
	SerialNumber *string `json:"serial_number,omitempty" query:"serial_number"`
	Activate     *bool   `json:"activate,omitempty" query:"activate"`
}

func (r *Filter) ToFilter() []string {
	filter := []string{
		"||PRELOAD||Slots",
		"||PRELOAD||Slots.Product",
	}

	if r.Limit != nil {
		filter = append(filter, fmt.Sprintf("||LIMIT||%d", *r.Limit))
	}

	if r.Offset != nil {
		filter = append(filter, fmt.Sprintf("||OFFSET||%d", *r.Offset))
	}

	if r.ID != nil {
		filter = append(filter, fmt.Sprintf("id||=||%d", *r.ID))
	}

	if r.Name != nil {
		filter = append(filter, fmt.Sprintf("name||=||%s", *r.Name))
	}

	if r.SerialNumber != nil {
		filter = append(filter, fmt.Sprintf("serial_number||=||%s", *r.SerialNumber))
	}

	if r.Activate != nil {
		filter = append(filter, fmt.Sprintf("activate||=||%v||bool", *r.Activate))
	}

	if r.SortBy != nil {
		val := strings.Split(*r.SortBy, ":")
		if len(val) == 1 || (val[1] != "asc" && val[1] != "desc") {
			filter = append(filter, fmt.Sprintf("%s||SORT||%s", val[0], "asc"))
		} else {
			filter = append(filter, fmt.Sprintf("%s||SORT||%s", val[0], val[1]))
		}
	} else {
		filter = append(filter, "id||SORT||asc")
	}

	return filter
}
