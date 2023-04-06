package auth_http

import "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth"

type restImpl struct {
	usecase auth.Usecase
}

func New(uc auth.Usecase) *restImpl {
	return &restImpl{uc}
}
