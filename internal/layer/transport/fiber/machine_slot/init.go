package machine_slot

import "vm-backend/internal/core/domain/machine"

type restImpl struct {
	usecase machine.SlotUsecase
}

func NewTransport(uc machine.SlotUsecase) machine.SlotTransport {
	return &restImpl{uc}
}
