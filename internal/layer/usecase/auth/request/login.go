package request

import "fmt"

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *Login) ToFilter() []string {
	return []string{
		fmt.Sprintf("username||=||%s", r.Username),
		"||PRELOAD||Role",
	}
}
