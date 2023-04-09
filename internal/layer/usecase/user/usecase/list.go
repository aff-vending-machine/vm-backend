package user_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) List(ctx context.Context, req *request.Filter) ([]response.User, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	users, err := uc.userRepo.FindMany(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "unable to find users")
	}

	return response.ToUserList(users), nil
}
