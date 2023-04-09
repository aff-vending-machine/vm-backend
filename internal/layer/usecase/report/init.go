package report

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

type usecaseImpl struct {
	machineRepo     repository.Machine
	machineSlotRepo repository.MachineSlot
	transactionRepo repository.Transaction
}

func New(m repository.Machine, s repository.MachineSlot, t repository.Transaction) *usecaseImpl {
	return &usecaseImpl{m, s, t}
}
