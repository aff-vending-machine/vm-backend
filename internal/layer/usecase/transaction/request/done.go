package request

import "fmt"

type Done struct {
	ID     uint   `json:"id" query:"id" validate:"required"`
	Caller string `json:"caller" query:"caller" validate:"required"`
}

func (r *Done) ToFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.ID),
	}
}
