package user_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Create(ctx context.Context, req *request.Create) (uint, error) {
	if v := validate.Struct(req); !v.Validate() {
		return 0, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	hashed, err := uc.passwordMgr.HashPassword(ctx, req.Password)
	if err != nil {
		return 0, errors.Wrap(err, "unable to hash password")
	}

	user := req.ToEntity(hashed)
	err = uc.userRepo.InsertOne(ctx, user)
	if err != nil {
		return 0, errors.Wrap(err, "unable to insert user")
	}

	return user.ID, nil
}
