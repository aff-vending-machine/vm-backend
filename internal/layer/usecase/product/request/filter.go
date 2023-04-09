package request

import (
	"fmt"
	"strings"
)

type Filter struct {
	Limit       *int    `json:"limit,omitempty" query:"limit"`
	Offset      *int    `json:"offset,omitempty" query:"offset"`
	SortBy      *string `json:"sort_by,omitempty" query:"sort_by"`
	ID          *uint   `json:"id,omitempty" query:"id"`
	Name        *string `json:"name,omitempty" query:"name"`
	SKU         *string `json:"sku,omitempty" query:"sku"`
	ProductType *string `json:"product_type,omitempty" query:"product_type"`
}

func (r *Filter) ToFilter() []string {
	filter := []string{}

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

	if r.SKU != nil {
		filter = append(filter, fmt.Sprintf("sku||=||%s", *r.SKU))
	}

	if r.ProductType != nil {
		filter = append(filter, fmt.Sprintf("product_type||=||%s", *r.ProductType))
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
