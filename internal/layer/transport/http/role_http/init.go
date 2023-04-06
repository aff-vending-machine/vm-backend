package role_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role"
)

type restImpl struct {
	usecase role.Usecase
}

func New(uc role.Usecase) *restImpl {
	return &restImpl{uc}
}
