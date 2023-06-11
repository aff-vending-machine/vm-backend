package response

import (
	"time"
)

type User struct {
	ID        uint        `json:"id"`
	BranchID  *uint       `json:"branch_id,omitempty"`
	Branch    *UserBranch `json:"branch,omitempty"`
	RoleID    uint        `json:"role_id"`
	Role      UserRole    `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Username  string      `json:"username"`
	Contact   string      `json:"contact"`
	CreatedBy string      `json:"created_by"`
	LastLogin *time.Time  `json:"last_login"`
}

type UserBranch struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	IsEnable bool   `json:"is_enable"`
}

type UserRole struct {
	Name        string           `json:"name"`
	Permissions []UserPermission `json:"permissions"`
}

type UserPermission struct {
	Scope string `json:"scope"`
	Level int    `json:"level"`
}
