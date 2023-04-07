package api

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/sync"
)

type RPC interface {
	GetMachine(context.Context, string) (*sync.Machine, error)
	GetSlot(context.Context, string) ([]sync.Slot, error)
	SetSlot(context.Context, string, []sync.Slot) error
	GetTransaction(context.Context, string) ([]sync.Transaction, error)
	ClearTransaction(context.Context, string, []uint) error
}
