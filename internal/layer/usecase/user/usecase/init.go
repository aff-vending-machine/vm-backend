package user_usecase

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/crypto"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

type usecaseImpl struct {
	userRepo    repository.User
	roleRepo    repository.Role
	passwordMgr crypto.Password
}

func New(u repository.User, r repository.Role, p crypto.Password) *usecaseImpl {
	return &usecaseImpl{u, r, p}
}
