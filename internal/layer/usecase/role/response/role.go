package response

import (
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Role struct {
	ID          uint         `json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

func ToRole(e *entity.Role) *Role {
	return &Role{
		ID:          e.ID,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		Name:        e.Name,
		Permissions: PermissionEntityToList(e.Permissions),
	}
}

func ToRoleList(roles []entity.Role) []Role {
	items := make([]Role, len(roles))
	for i, role := range roles {
		items[i] = *ToRole(&role)
	}

	return items
}
