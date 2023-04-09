package request

import (
	"fmt"
)

type ResetPassword struct {
	ID    uint `json:"id" query:"id" validate:"required"`
	Level int  `json:"level" query:"level" validate:"required"`
}

func (r *ResetPassword) ToFilter() []string {
	return []string{
		fmt.Sprintf("id||=||%d", r.ID),
	}
}

func (r *ResetPassword) ToJson(hashed string) map[string]interface{} {
	return map[string]interface{}{
		"password": hashed,
	}
}
