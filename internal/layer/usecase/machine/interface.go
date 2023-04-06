package machine

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine/response"
)

type Usecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Get(context.Context, *request.Get) (*response.Machine, error)
	Create(context.Context, *request.Create) (uint, error)
	Delete(context.Context, *request.Delete) error
	List(context.Context, *request.Filter) ([]response.Machine, error)
	Update(context.Context, *request.Update) error
}
