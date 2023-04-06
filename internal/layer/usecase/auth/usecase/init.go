package auth_usecase

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/crypto"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

type usecase struct {
	roleRepo    repository.Role
	userRepo    repository.User
	passwordMgr crypto.Password
	tokenMgr    crypto.Token
}

func New(r repository.Role, u repository.User, p crypto.Password, t crypto.Token) *usecase {
	return &usecase{r, u, p, t}
}
