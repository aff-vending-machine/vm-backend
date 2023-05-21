package request

import (
	"vm-backend/pkg/conv"
	"vm-backend/pkg/db"
)

type Update struct {
	ID       uint    `json:"id" query:"id" validate:"required"`
	Name     *string `json:"name,omitempty"`
	IsEnable *bool   `json:"is_enable,omitempty"`
}

func (r *Update) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID)
}

func (r *Update) ToUpdate() map[string]interface{} {
	result, _ := conv.StructToMap(r)
	delete(result, "id")
	return result
}
