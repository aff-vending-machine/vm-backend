package machine_usecase

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

type usecaseImpl struct {
	machineRepo     repository.Machine
	machineSlotRepo repository.MachineSlot
}

func New(m repository.Machine, s repository.MachineSlot) *usecaseImpl {
	return &usecaseImpl{m, s}
}
