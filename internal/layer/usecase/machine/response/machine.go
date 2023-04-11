package response

import (
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/response"
)

type Machine struct {
	ID                  uint                   `json:"id"`
	Name                string                 `json:"name"`
	SerialNumber        string                 `json:"serial_number"`
	Location            string                 `json:"location"`
	Type                string                 `json:"type"`
	Vendor              string                 `json:"vendor"`
	Slots               []response.MachineSlot `json:"slots"`
	Status              string                 `json:"status"`
	SyncMachineTime     *time.Time             `json:"sync_machine_time"`
	SyncSlotTime        *time.Time             `json:"sync_slot_time"`
	SyncTransactionTime *time.Time             `json:"sync_transaction_time"`
}

func ToMachine(e *entity.Machine) *Machine {
	return &Machine{
		ID:                  e.ID,
		Name:                e.Name,
		SerialNumber:        e.SerialNumber,
		Location:            e.Location,
		Type:                e.Type,
		Vendor:              e.Vendor,
		Slots:               response.ToMachineSlotList(e.Slots),
		Status:              e.Status,
		SyncMachineTime:     e.SyncMachineTime,
		SyncSlotTime:        e.SyncSlotTime,
		SyncTransactionTime: e.SyncTransactionTime,
	}
}

func MachineEntityToList(es []entity.Machine) []Machine {
	items := make([]Machine, len(es))
	for i, e := range es {
		items[i] = *ToMachine(&e)
	}

	return items
}
