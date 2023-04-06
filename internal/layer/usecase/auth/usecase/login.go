package auth_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecase) Login(ctx context.Context, req *request.Login) (*response.AuthResult, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	user, err := uc.userRepo.FindOne(ctx, req.ToFilter())
	if err != nil {
		return nil, errors.Wrap(err, "unable to find one user")
	}

	ok, err := uc.passwordMgr.Compare(ctx, user.Password, req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "unable to compare password")
	}
	if !ok {
		return nil, errors.Wrap(err, "compare password is mismatched")
	}

	token, err := uc.generateToken(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "generate token failed")
	}

	return token, nil
}
