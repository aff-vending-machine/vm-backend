package account

import (
	"strings"
	"time"
	"vm-backend/internal/core/infra/strorage/postgresql/service"
)

type Permission struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	RoleID    uint      `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Scope     string    `json:"scope"`
	AltScope  string    `json:"alt_scope"`
	Level     int       `json:"level"`
}

func (e Permission) TableName() string {
	return "account_permissions"
}

func (e Permission) HasScope(s string) bool {
	scope := strings.ToLower(s)
	return strings.ToLower(e.Scope) == scope || strings.ToLower(e.AltScope) == scope
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
