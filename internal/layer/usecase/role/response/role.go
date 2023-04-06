package response

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Role struct {
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

func RoleEntityToView(role *entity.Role) *Role {
	return &Role{
		Name:        role.Name,
		Permissions: PermissionEntityToList(role.Permissions),
	}
}

func RoleDomainToList(roles []entity.Role) []Role {
	items := make([]Role, len(roles))
	for i, role := range roles {
		items[i] = *RoleEntityToView(&role)
	}

	return items
}
