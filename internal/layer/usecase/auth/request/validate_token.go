package request

import "fmt"

type ValidateToken struct {
	Token string `json:"token" validate:"required"`
}

func (r *ValidateToken) ToFilter(id uint) []string {
	return []string{
		fmt.Sprintf("id:=:%d", id),
		":PRELOAD:Role",
	}
}
