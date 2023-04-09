package response

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Permission struct {
	Scope string `json:"scope"`
	Level int    `json:"level"`
}

func ToPermission(p *entity.Permission) *Permission {
	return &Permission{
		Scope: p.Scope,
		Level: p.Level,
	}
}

func ToPermissionList(ps []entity.Permission) []Permission {
	items := make([]Permission, len(ps))
	for i, p := range ps {
		items[i] = *ToPermission(&p)
	}

	return items
}
