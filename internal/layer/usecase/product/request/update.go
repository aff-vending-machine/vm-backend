package request

import (
	"encoding/json"
	"fmt"
)

type Update struct {
	ID       uint     `json:"id" query:"id" validate:"required"`
	Name     *string  `json:"name,omitempty"`
	Type     *string  `json:"type,omitempty"`
	ImageURL *string  `json:"image_url,omitempty"`
	Price    *float64 `json:"price,omitempty"`
}

func (r *Update) ToFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.ID),
	}
}

func (r *Update) ToJson() map[string]interface{} {
	var data map[string]interface{}

	b, _ := json.Marshal(r)
	json.Unmarshal(b, &data)

	delete(data, "id")
	return data
}
