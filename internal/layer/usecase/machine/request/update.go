package request

import (
	"vm-backend/pkg/conv"
	"vm-backend/pkg/db"
)

type Update struct {
	ID           uint    `json:"id" query:"id" validate:"required"`
	BranchID     *uint   `json:"branch_id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Location     *string `json:"location,omitempty"`
	Type         *string `json:"type,omitempty"`
	Vendor       *string `json:"vendor,omitempty"`
}

func (r *Update) ToQuery() *db.Query {
	return db.NewQuery().AddWhere("id = ?", r.ID)
}

func (r *Update) ToUpdate() map[string]interface{} {
	result, _ := conv.StructToMap(r)
	delete(result, "id")
	return result
}
