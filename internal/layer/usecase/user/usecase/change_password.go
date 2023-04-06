package user_usecase

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) ChangePassword(ctx context.Context, req *request.ChangePassword) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	based, err := uc.userRepo.FindOne(ctx, filter)
	if err != nil {
		return errors.Wrapf(err, "unable to find user %d", req.ID)
	}

	ok, err := uc.passwordMgr.Compare(ctx, based.Password, req.OldPassword)
	if err != nil {
		return errors.Wrap(err, "unable to compare password")
	}
	if !ok {
		return fmt.Errorf("password is not match")
	}

	hashed, err := uc.passwordMgr.HashPassword(ctx, req.NewPassword)
	if err != nil {
		return errors.Wrap(err, "unable to hash password")
	}

	data := req.ToJson(hashed)
	_, err = uc.userRepo.UpdateMany(ctx, filter, data)
	if err != nil {
		return errors.Wrapf(err, "unable to update user %d failed", req.ID)
	}

	return nil
}
