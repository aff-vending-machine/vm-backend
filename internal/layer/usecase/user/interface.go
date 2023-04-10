package user

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/response"
)

type Usecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Get(context.Context, *request.Get) (*response.User, error)
	List(context.Context, *request.Filter) ([]response.User, error)
	Create(context.Context, *request.Create) (uint, error)
	ChangeRole(context.Context, *request.ChangeRole) error
	ChangePassword(context.Context, *request.ChangePassword) error
	ResetPassword(context.Context, *request.ResetPassword) error
	Delete(context.Context, *request.Delete) error
}
