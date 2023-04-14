package usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/sync/request"
)

type Sync interface {
	FetchMachine(context.Context, *request.Sync) error
	PushMachine(context.Context, *request.Sync) error
	FetchSlots(context.Context, *request.Sync) error
	PushSlots(context.Context, *request.Sync) error
	PullTransactions(context.Context, *request.Sync) error
	RegisterMachine(context.Context, *request.RegisterMachine) error
}
