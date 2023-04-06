package auth_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecase) ValidateRefreshToken(ctx context.Context, req *request.ValidateToken) (*response.AuthResult, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	claims, err := uc.tokenMgr.ValidateRefreshToken(ctx, req.Token)
	if err != nil {
		return nil, errors.Wrap(err, "unable to validate refresh token")
	}

	filter := req.ToFilter(claims.ID)
	user, err := uc.userRepo.FindOne(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "unable to find user")
	}

	token, err := uc.generateToken(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "unable to generate token")
	}

	return token, nil
}
