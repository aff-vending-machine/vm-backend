package auth

import "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth"

type httpImpl struct {
	usecase auth.Usecase
}

func New(uc auth.Usecase) *httpImpl {
	return &httpImpl{uc}
}
