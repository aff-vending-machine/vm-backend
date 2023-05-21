package account

import (
	"context"
	"vm-backend/internal/layer/usecase/account/request"
	"vm-backend/internal/layer/usecase/account/response"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecase) Login(ctx context.Context, req *request.Login) (*response.AuthResult, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToUserQuery()
	user, err := uc.userRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find user")
		return nil, errors.Wrap(err, "unable to find user")
	}

	ok, err := uc.passwordMgr.Compare(ctx, user.Password, req.Password)
	if err != nil {
		log.Error().Err(err).Interface("password", req.Password).Msg("compare password failed")
		return nil, errors.Wrap(err, "unable to compare password")
	}
	if !ok {
		log.Error().Interface("password", req.Password).Msg("password not match")
		return nil, errors.New("password not match")
	}

	token, err := uc.generateToken(ctx, user)
	if err != nil {
		log.Error().Interface("user", user).Err(err).Msg("generate token failed")
		return nil, errors.Wrap(err, "unable to generate token")
	}

	return token, nil
}
