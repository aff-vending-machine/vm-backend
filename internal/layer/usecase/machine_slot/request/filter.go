package request

import (
	"fmt"
)

type Filter struct {
	MachineID uint    `json:"machine_id" query:"machine_id"`
	ID        *uint   `json:"id,omitempty" query:"id"`
	ProductID *uint   `json:"product_id,omitempty" query:"product_id"`
	Code      *string `json:"code,omitempty" query:"code"`
	Stock     *int    `json:"stock,omitempty" query:"stock"`
	Capacity  *int    `json:"capacity,omitempty" query:"capacity"`
}

func (r *Filter) ToFilter() []string {
	filter := []string{
		fmt.Sprintf("machine_id||=||%d", r.MachineID),
		"||PRELOAD||Product",
		"id||SORT||asc",
	}

	if r.ID != nil {
		filter = append(filter, fmt.Sprintf("id||=||%d", *r.ID))
	}

	if r.ProductID != nil {
		filter = append(filter, fmt.Sprintf("product_id||=||%d", *r.ProductID))
	}

	if r.Code != nil {
		filter = append(filter, fmt.Sprintf("code||=||%s", *r.Code))
	}

	if r.Stock != nil {
		filter = append(filter, fmt.Sprintf("stock||=||%d", *r.Stock))
	}

	if r.Capacity != nil {
		filter = append(filter, fmt.Sprintf("capacity||=||%d", *r.Capacity))
	}

	return filter
}
