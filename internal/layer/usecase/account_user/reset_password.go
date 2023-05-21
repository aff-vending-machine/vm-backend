package account_user

import (
	"context"

	"vm-backend/internal/layer/usecase/account_user/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) ResetPassword(ctx context.Context, req *request.ResetPassword) error {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return errors.Wrap(err, "unable to validate request")
	}

	pwd := "00000000"
	hashed, err := uc.passwordMgr.HashPassword(ctx, pwd)
	if err != nil {
		log.Error().Err(err).Msg("unable to hash password")
		return errors.Wrap(err, "unable to hash password")
	}

	query := req.ToQuery()
	data := req.ToUpdate(hashed)
	_, err = uc.userRepo.Update(ctx, query, data)
	if err != nil {
		log.Error().Err(err).Msg("unable to update user")
		return errors.Wrap(err, "unable to update user")
	}

	return nil
}
