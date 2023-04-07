package usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/sync/request"
)

type Sync interface {
	GetMachine(context.Context, *request.Sync) error
	GetSlot(context.Context, *request.Sync) error
	SetSlot(context.Context, *request.Sync) error
	GetTransaction(context.Context, *request.Sync) error
	RegisterMachine(context.Context, *request.RegisterMachine) error
}
