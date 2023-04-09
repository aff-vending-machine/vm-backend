package entity

import (
	"time"
)

type User struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Username  string     `json:"username" gorm:"uniqueIndex"`
	Password  string     `json:"-"`
	RoleID    uint       `json:"role_id"`
	Role      Role       `json:"role"`
	CreatedBy string     `json:"created_by"`
	LastLogin *time.Time `json:"last_login"`
}

func (e User) TableName() string {
	return "users"
}

func (e User) HasRole(name string) bool {
	return e.Role.Name == name
}
