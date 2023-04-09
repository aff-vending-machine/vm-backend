package request

import (
	"fmt"
)

type Delete struct {
	ID uint `json:"id" query:"id" validate:"required"`
}

func (r *Delete) ToFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.ID),
	}
}
