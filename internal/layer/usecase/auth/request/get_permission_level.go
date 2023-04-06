package request

import "fmt"

type GetPermissionLevel struct {
	UserID uint   `json:"user_id" query:"id" validate:"required"`
	Scope  string `json:"scope" validate:"required"`
}

func (r *GetPermissionLevel) ToUserFilter() []string {
	return []string{
		fmt.Sprintf("id:=:%d", r.UserID),
		":PRELOAD:Role",
	}
}

func (r *GetPermissionLevel) ToRoleFilter(roleID uint) []string {
	return []string{
		fmt.Sprintf("id:=:%d", roleID),
		":PRELOAD:Permissions",
	}
}
