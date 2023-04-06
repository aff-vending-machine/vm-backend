package request

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Create struct {
	Name        string       `json:"name" validate:"required"`
	Permissions []Permission `json:"permissions" validate:"required"`
}

func (r *Create) ToEntity() *entity.Role {
	return &entity.Role{
		Name:        r.Name,
		Permissions: PermissionListToEntity(r.Permissions),
	}
}
