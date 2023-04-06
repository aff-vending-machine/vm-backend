package request

import (
	"fmt"
)

type ChangePassword struct {
	ID          uint   `json:"id" query:"id" validate:"required"`
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

func (r *ChangePassword) ToFilter() []string {
	return []string{
		fmt.Sprintf("id:=:%d", r.ID),
	}
}

func (r *ChangePassword) ToJson(hashed string) map[string]interface{} {
	return map[string]interface{}{
		"password": hashed,
	}
}
