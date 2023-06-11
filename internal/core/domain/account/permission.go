package account

import (
	"time"
	"vm-backend/internal/core/infra/strorage/postgresql/service"
)

type Permission struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	RoleID    uint      `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Scope     string    `json:"scope"`
	Level     int       `json:"level"`
}

func (e Permission) TableName() string {
	return "account_permissions"
}

func (e Permission) HasScope(scope string) bool {
	return e.Scope == scope
}

type PermissionRepository interface {
	service.Repository[Permission]
}

const (
	Forbidder = 0
	Viewer    = 1
	Owner     = 2
	Editor    = 3
	Admin     = 4
)
