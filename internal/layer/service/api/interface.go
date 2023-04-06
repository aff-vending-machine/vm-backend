package api

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/model"
)

type RPC interface {
	MachineGet(context.Context, string) (*model.Machine, error)
	SlotGet(context.Context, string) ([]model.Slot, error)
	SlotSet(context.Context, string, []model.Slot) error
}
