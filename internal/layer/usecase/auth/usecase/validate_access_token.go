package auth_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecase) ValidateAccessToken(ctx context.Context, req *request.ValidateToken) (*response.Claims, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	claims, err := uc.tokenMgr.ValidateAccessToken(ctx, req.Token)
	if err != nil {
		return nil, errors.Wrap(err, "unable to validate access token")
	}

	return response.ToToken(claims), nil
}
