package account_user

import (
	"context"
	"fmt"
	"vm-backend/internal/layer/usecase/account_user/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) ChangePassword(ctx context.Context, req *request.ChangePassword) error {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	based, err := uc.userRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find user")
		return errors.Wrapf(err, "unable to find user %d", req.ID)
	}

	ok, err := uc.passwordMgr.Compare(ctx, based.Password, req.OldPassword)
	if err != nil {
		log.Error().Err(err).Interface("password", req.OldPassword).Msg("unable to compare password")
		return errors.Wrap(err, "unable to compare password")
	}
	if !ok {
		log.Error().Interface("password", req.OldPassword).Msg("password is not match")
		return fmt.Errorf("password is not match")
	}

	hashed, err := uc.passwordMgr.HashPassword(ctx, req.NewPassword)
	if err != nil {
		log.Error().Err(err).Interface("password", req.NewPassword).Msg("unable to hash password")
		return errors.Wrap(err, "unable to hash password")
	}

	data := req.ToUpdate(hashed)
	_, err = uc.userRepo.Update(ctx, query, data)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Interface("data", data).Msg("unable to update user")
		return errors.Wrapf(err, "unable to update user %d", req.ID)
	}

	return nil
}
