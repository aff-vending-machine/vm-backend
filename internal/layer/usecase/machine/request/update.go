package request

import (
	"vm-backend/pkg/helpers/conv"
	"vm-backend/pkg/helpers/db"
)

type Update struct {
	ID               uint    `json:"id" query:"id" validate:"required"`
	BranchIDForQuery *uint   `json:"-" query:"branch_id"`
	BranchID         *uint   `json:"branch_id,omitempty"`
	Name             *string `json:"name,omitempty"`
	Location         *string `json:"location,omitempty"`
	Type             *string `json:"type,omitempty"`
	Vendor           *string `json:"vendor,omitempty"`
}

func (r *Update) ToQuery() *db.Query {
	return db.NewQuery().
		Where("id = ?", r.ID).
		WhereIf("branch_id = ?", r.BranchIDForQuery)
}

func (r *Update) ToUpdate() map[string]interface{} {
	result, _ := conv.StructToMap(r)
	delete(result, "id")
	return result
}
