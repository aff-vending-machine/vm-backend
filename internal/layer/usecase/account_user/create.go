package account_user

import (
	"context"
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/layer/usecase/account_user/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Create(ctx context.Context, req *request.Create) (uint, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return 0, errors.Wrap(err, "unable to validate request")
	}

	hashed, err := uc.passwordMgr.HashPassword(ctx, req.Password)
	if err != nil {
		log.Error().Err(err).Msg("unable to hash password")
		return 0, errors.Wrap(err, "unable to hash password")
	}

	user := makeUser(req, hashed)
	_, err = uc.userRepo.Create(ctx, user)
	if err != nil {
		log.Error().Err(err).Msg("unable to create user")
		return 0, errors.Wrap(err, "unable to create user")
	}

	return user.ID, nil
}

func makeUser(req *request.Create, hashed string) *account.User {
	return &account.User{
		Username:  req.Username,
		Password:  hashed,
		CreatedBy: req.CreatedBy,
		RoleID:    req.RoleID,
	}
}
