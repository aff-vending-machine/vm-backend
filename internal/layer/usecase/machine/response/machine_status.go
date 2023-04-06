package response

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type MachineStatus struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}

func MachineEntityToStatusView(e *entity.Machine) *MachineStatus {
	return &MachineStatus{
		ID:     e.ID,
		Status: e.Status,
	}
}
