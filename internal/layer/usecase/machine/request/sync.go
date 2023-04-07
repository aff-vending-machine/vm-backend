package request

import (
	"fmt"
)

type Sync struct {
	ID uint `json:"id" query:"id" validate:"required"`
}

func (r *Sync) ToFilter() []string {
	return []string{
		fmt.Sprintf("id:=:%d", r.ID),
	}
}
