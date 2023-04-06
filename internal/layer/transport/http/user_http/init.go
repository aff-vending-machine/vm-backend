package user_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user"
)

type restImpl struct {
	usecase user.Usecase
}

func New(uc user.Usecase) *restImpl {
	return &restImpl{uc}
}
