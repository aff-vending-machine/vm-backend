package machine_slot_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot"
)

type restImpl struct {
	usecase machine_slot.Usecase
}

func New(uc machine_slot.Usecase) *restImpl {
	return &restImpl{uc}
}
