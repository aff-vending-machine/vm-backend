package request

import (
	"fmt"
)

type Get struct {
	ID uint `json:"id" query:"id" validate:"required"`
}

func (r *Get) ToFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.ID),
		"||PRELOAD||Slots",
		"||PRELOAD||Slots.Product",
	}
}
