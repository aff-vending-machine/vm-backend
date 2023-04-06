package request

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Create struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleID   uint   `json:"role_id" validate:"required"`
}

func (r *Create) ToEntity(hash string) *entity.User {
	return &entity.User{
		Username: r.Username,
		Password: hash,
		RoleID:   r.RoleID,
	}
}
