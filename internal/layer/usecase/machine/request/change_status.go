package request

import (
	"fmt"
)

type ChangeStatus struct {
	ID     uint   `json:"id" query:"id" validate:"required"`
	Status string `json:"status"  validate:"required"`
}

func (r *ChangeStatus) ToFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.ID),
	}
}

func (r *ChangeStatus) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"status": r.Status,
	}
}
