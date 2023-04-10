package machine

import "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine"

type restImpl struct {
	usecase machine.Usecase
}

func New(uc machine.Usecase) *restImpl {
	return &restImpl{uc}
}
