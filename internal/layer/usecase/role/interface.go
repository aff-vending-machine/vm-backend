package role

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role/response"
)

type Usecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Get(context.Context, *request.Get) (*response.Role, error)
	List(context.Context, *request.Filter) ([]response.Role, error)
	Create(context.Context, *request.Create) (uint, error)
	Update(context.Context, *request.Update) error
	Delete(context.Context, *request.Delete) error
}
