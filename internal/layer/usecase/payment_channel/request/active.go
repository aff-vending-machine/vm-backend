package request

import (
	"encoding/json"
	"fmt"
)

type Active struct {
	ID     uint `json:"id" query:"id" validate:"required"`
	Active bool `json:"active" validate:"required"`
}

func (r *Active) ToFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.ID),
	}
}

func (r *Active) ToJson() map[string]interface{} {
	var data map[string]interface{}

	b, _ := json.Marshal(r)
	json.Unmarshal(b, &data)

	delete(data, "id")
	return data
}
