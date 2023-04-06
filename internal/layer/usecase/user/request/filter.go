package request

import (
	"fmt"
)

type Filter struct {
	Limit    *int    `json:"limit,omitempty" query:"limit"`
	Offset   *int    `json:"offset,omitempty" query:"offset"`
	ID       *uint   `json:"id,omitempty" query:"id"`
	Username *string `json:"username,omitempty" query:"username"`
}

func (r *Filter) ToFilter() []string {
	filter := []string{
		":PRELOAD:Role",
	}

	if r.Limit != nil {
		filter = append(filter, fmt.Sprintf(":LIMIT:%d", *r.Limit))
	}

	if r.Offset != nil {
		filter = append(filter, fmt.Sprintf(":OFFSET:%d", *r.Offset))
	}

	if r.ID != nil {
		filter = append(filter, fmt.Sprintf("id:=:%d", *r.ID))
	}

	if r.Username != nil {
		filter = append(filter, fmt.Sprintf("username:=:%s", *r.Username))
	}

	return filter
}
