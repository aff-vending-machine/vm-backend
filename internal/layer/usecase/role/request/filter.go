package request

import (
	"fmt"
)

type Filter struct {
	Limit  *int    `json:"limit,omitempty" query:"limit"`
	Offset *int    `json:"offset,omitempty" query:"offset"`
	ID     *uint   `json:"id,omitempty" query:"id"`
	Name   *string `json:"name,omitempty" query:"name"`
}

func (r *Filter) ToFilter() []string {
	filter := []string{}

	if r.Limit != nil {
		filter = append(filter, fmt.Sprintf(":LIMIT:%d", *r.Limit))
	}

	if r.Offset != nil {
		filter = append(filter, fmt.Sprintf(":OFFSET:%d", *r.Offset))
	}

	if r.ID != nil {
		filter = append(filter, fmt.Sprintf("id:=:%d", *r.ID))
	}

	if r.Name != nil {
		filter = append(filter, fmt.Sprintf("name:=:%s", *r.Name))
	}

	return filter
}
