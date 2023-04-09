package request

import (
	"fmt"
	"strings"
)

type Filter struct {
	Limit    *int    `json:"limit,omitempty" query:"limit"`
	Offset   *int    `json:"offset,omitempty" query:"offset"`
	SortBy   *string `json:"sort_by,omitempty" query:"sort_by"`
	ID       *uint   `json:"id,omitempty" query:"id"`
	Username *string `json:"username,omitempty" query:"username"`
}

func (r *Filter) ToFilter() []string {
	filter := []string{
		"||PRELOAD||Role",
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

	if r.Username != nil {
		filter = append(filter, fmt.Sprintf("username||=||%s", *r.Username))
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
