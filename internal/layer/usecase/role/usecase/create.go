package role_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Create(ctx context.Context, req *request.Create) (uint, error) {
	if v := validate.Struct(req); !v.Validate() {
		return 0, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	role := req.ToEntity()
	err := uc.roleRepo.InsertOne(ctx, role)
	if err != nil {
		return 0, errors.Wrap(err, "unable to insert role")
	}

	return role.ID, nil
}
