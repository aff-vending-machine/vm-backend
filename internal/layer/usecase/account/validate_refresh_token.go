package account

import (
	"context"
	"vm-backend/internal/layer/usecase/account/request"
	"vm-backend/internal/layer/usecase/account/response"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecase) ValidateRefreshToken(ctx context.Context, req *request.ValidateToken) (*response.AuthResult, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	claims, err := uc.tokenMgr.ValidateRefreshToken(ctx, req.Token)
	if err != nil {
		log.Error().Err(err).Interface("request", req).Msg("unable to validate token")
		return nil, errors.Wrap(err, "unable to validate token")
	}

	query := req.ToUserQuery(claims.ID)
	user, err := uc.userRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("request", req).Msg("unable to find user")
		return nil, errors.Wrap(err, "unable to find user")
	}

	token, err := uc.generateToken(ctx, user)
	if err != nil {
		log.Error().Err(err).Interface("request", req).Msg("unable to generate token")
		return nil, errors.Wrap(err, "unable to generate token")
	}

	return token, nil
}
