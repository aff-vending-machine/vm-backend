package request

import "fmt"

type Cancel struct {
	ID     uint   `json:"id" query:"id" validate:"required"`
	Caller string `json:"caller" query:"caller" validate:"required"`
}

func (r *Cancel) ToFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.ID),
	}
}
