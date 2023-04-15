package machine_slot

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/response"
)

type Usecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Create(context.Context, *request.Create) (uint, error)
	Delete(context.Context, *request.Delete) error
	Get(context.Context, *request.Get) (*response.MachineSlot, error)
	List(context.Context, *request.Filter) ([]response.MachineSlot, error)
	Update(context.Context, *request.Update) error
	BulkUpdate(context.Context, *request.BulkUpdate) error
}
