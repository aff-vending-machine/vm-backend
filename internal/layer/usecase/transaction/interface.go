package transaction

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/response"
)

type Usecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Get(context.Context, *request.Get) (*response.Transaction, error)
	List(context.Context, *request.Filter) ([]response.Transaction, error)
	Create(context.Context, *request.Create) (uint, error)
	Update(context.Context, *request.Update) error
	Delete(context.Context, *request.Delete) error
}
