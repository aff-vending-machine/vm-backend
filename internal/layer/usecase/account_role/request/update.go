package request

import (
	"vm-backend/pkg/conv"
	"vm-backend/pkg/helpers/db"
)

type Update struct {
	ID          uint         `json:"id" query:"id" validate:"required"`
	Name        string       `json:"name,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
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
