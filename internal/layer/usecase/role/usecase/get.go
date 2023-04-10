package role_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Get(ctx context.Context, req *request.Get) (*response.Role, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	role, err := uc.roleRepo.FindOne(ctx, filter)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find role %d", req.ID)
	}

	return response.ToRole(role), nil
}
