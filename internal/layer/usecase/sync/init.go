package sync

import (
	"context"
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

type usecaseImpl struct {
	rpcAPI          api.RPC
	machineRepo     repository.Machine
	machineSlotRepo repository.MachineSlot
	productRepo     repository.Product
	transactionRepo repository.Transaction
}

func New(r api.RPC, m repository.Machine, s repository.MachineSlot, p repository.Product, t repository.Transaction) *usecaseImpl {
	return &usecaseImpl{r, m, s, p, t}
}

func makeCodeFilter(machineID uint, code string) []string {
	return []string{
		fmt.Sprintf("machine_id:=:%d", machineID),
		fmt.Sprintf("code:=:%s", code),
	}
}

func (uc *usecaseImpl) updateMachineStatus(ctx context.Context, filter []string, status string) {
	data := map[string]interface{}{
		"status": status,
	}

	switch status {
	case "active":
		data["last_active_time"] = time.Now()
	case "maintenance":
		data["last_maintenance_time"] = time.Now()
	}

	uc.machineRepo.UpdateMany(ctx, filter, data)
}
