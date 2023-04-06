package user_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) ResetPassword(ctx context.Context, req *request.ResetPassword) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	pwd := "00000000"
	hashed, err := uc.passwordMgr.HashPassword(ctx, pwd)
	if err != nil {
		return errors.Wrap(err, "unable to hash password")
	}

	filter := req.ToFilter()
	data := req.ToJson(hashed)
	_, err = uc.userRepo.UpdateMany(ctx, filter, data)
	if err != nil {
		return errors.Wrapf(err, "unable to update user %d", req.ID)
	}

	return nil
}
