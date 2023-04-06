package machine_slot_usecase

import (
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

type usecaseImpl struct {
	rpcAPI          api.RPC
	machineRepo     repository.Machine
	machineSlotRepo repository.MachineSlot
	productRepo     repository.Product
}

func New(r api.RPC, m repository.Machine, s repository.MachineSlot, p repository.Product) *usecaseImpl {
	return &usecaseImpl{r, m, s, p}
}

func makeCodeFilter(machineID uint, code string) []string {
	return []string{
		fmt.Sprintf("machine_id:=:%d", machineID),
		fmt.Sprintf("code:=:%s", code),
	}
}
