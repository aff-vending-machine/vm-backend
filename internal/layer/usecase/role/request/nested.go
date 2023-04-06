package request

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Permission struct {
	Scope string `json:"scope"`
	Level int    `json:"level"`
}

func (r *Permission) ToEntity() *entity.Permission {
	return &entity.Permission{
		Scope: r.Scope,
		Level: r.Level,
	}
}

func PermissionListToEntity(permissions []Permission) []entity.Permission {
	items := make([]entity.Permission, len(permissions))
	for i, permission := range permissions {
		items[i] = *permission.ToEntity()
	}

	return items
}
