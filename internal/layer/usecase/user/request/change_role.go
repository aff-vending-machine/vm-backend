package request

import (
	"fmt"
)

type ChangeRole struct {
	ID     uint `json:"id" query:"id" validate:"required"`
	RoleID uint `json:"role_id" validate:"required"`
}

func (r *ChangeRole) ToFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.ID),
	}
}

func (r *ChangeRole) ToRoleFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.RoleID),
	}
}

func (r *ChangeRole) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"role_id": r.RoleID,
	}
}
