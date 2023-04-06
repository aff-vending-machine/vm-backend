package payment_channel

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel/response"
)

type Usecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Get(context.Context, *request.Get) (*response.PaymentChannel, error)
	List(context.Context, *request.Filter) ([]response.PaymentChannel, error)
	Create(context.Context, *request.Create) (uint, error)
	Update(context.Context, *request.Update) error
	Delete(context.Context, *request.Delete) error
	Active(context.Context, *request.Active) error
}
