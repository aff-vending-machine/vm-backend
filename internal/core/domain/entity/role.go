package entity

import (
	"time"
)

type Role struct {
	ID          uint         `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Name        string       `json:"name" gorm:"uniqueIndex"`
	Permissions []Permission `json:"permissions"`
}

func (e Role) TableName() string {
	return "roles"
}

func (e Role) HasPermission(scope string) int {
	for _, permission := range e.Permissions {
		if permission.Scope == scope {
			return permission.Level
		}
	}

	return 0
}
