package request

import (
	"encoding/json"
	"fmt"
)

type Update struct {
	ID          uint    `json:"id" query:"id" validate:"required"`
	Name        *string `json:"name"`
	Channel     *string `json:"channel"`
	Reference   *[]byte `json:"reference"`
	Secret      *[]byte `json:"secret"`
	DocumentURL *string `json:"document_url"`
	LogoURL     *string `json:"logo_url"`
	Vendor      *string `json:"vendor"`
	Active      *bool   `json:"active"`
}

func (r *Update) ToFilter() []string {
	return []string{
		fmt.Sprintf("id:=:%d", r.ID),
	}
}

func (r *Update) ToJson() map[string]interface{} {
	var data map[string]interface{}

	b, _ := json.Marshal(r)
	json.Unmarshal(b, &data)

	delete(data, "id")
	return data
}
