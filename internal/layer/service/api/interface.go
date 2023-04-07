package api

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/model"
)

type RPC interface {
	GetMachine(context.Context, string) (*model.Machine, error)
	GetSlot(context.Context, string) ([]model.Slot, error)
	SetSlot(context.Context, string, []model.Slot) error
}
