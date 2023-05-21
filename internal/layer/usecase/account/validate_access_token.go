package account

import (
	"context"
	"vm-backend/internal/layer/usecase/account/request"
	"vm-backend/internal/layer/usecase/account/response"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecase) ValidateAccessToken(ctx context.Context, req *request.ValidateToken) (*response.Claims, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	claims, err := uc.tokenMgr.ValidateAccessToken(ctx, req.Token)
	if err != nil {
		log.Error().Err(err).Str("token", req.Token).Msg("unable to validate token")
		return nil, errors.Wrap(err, "unable to validate token")
	}

	return &response.Claims{
		UserID:   claims.ID,
		Username: claims.Name,
		Role:     claims.Role,
		Type:     claims.Type,
	}, nil
}
