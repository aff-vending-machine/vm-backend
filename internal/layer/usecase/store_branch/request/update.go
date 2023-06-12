package request

import (
	"vm-backend/pkg/helpers/conv"
	"vm-backend/pkg/helpers/db"
)

type Update struct {
	ID       uint    `json:"id" query:"id" validate:"required"`
	Name     *string `json:"name,omitempty"`
	IsEnable *bool   `json:"is_enable,omitempty"`
}

func (r *Update) ToQuery() *db.Query {
	return db.NewQuery().
		Where("id = ?", r.ID)
}

func (r *Update) ToUpdate() map[string]interface{} {
	result, _ := conv.ToMap(r)
	delete(result, "id")
	return result
}
