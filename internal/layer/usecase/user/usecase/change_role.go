package user_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) ChangeRole(ctx context.Context, req *request.ChangeRole) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToRoleFilter()
	_, err := uc.roleRepo.FindOne(ctx, filter)
	if err != nil {
		return errors.Wrapf(err, "unable to find role %d", req.RoleID)
	}

	filter = req.ToFilter()
	data := req.ToJson()
	_, err = uc.userRepo.UpdateMany(ctx, filter, data)
	if err != nil {
		return errors.Wrapf(err, "unable to update user %d", req.ID)
	}

	return nil
}
