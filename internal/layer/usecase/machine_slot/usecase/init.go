package machine_slot_usecase

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

type usecaseImpl struct {
	machineRepo     repository.Machine
	machineSlotRepo repository.MachineSlot
	productRepo     repository.Product
}

func New(m repository.Machine, s repository.MachineSlot, p repository.Product) *usecaseImpl {
	return &usecaseImpl{m, s, p}
}
