package machine_usecase

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

type usecaseImpl struct {
	rpcAPI          api.RPC
	machineRepo     repository.Machine
	machineSlotRepo repository.MachineSlot
}

func New(r api.RPC, m repository.Machine, s repository.MachineSlot) *usecaseImpl {
	return &usecaseImpl{r, m, s}
}
