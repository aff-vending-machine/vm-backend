package account_user

import (
	"context"
	"vm-backend/internal/layer/usecase/account_user/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) ChangeRole(ctx context.Context, req *request.ChangeRole) error {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return errors.Wrap(err, "unable to validate request")
	}

	query := req.ToRoleQuery()
	_, err := uc.roleRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find role")
		return errors.Wrapf(err, "unable to find role %d", req.RoleID)
	}

	query = req.ToQuery()
	data := req.ToUpdate()
	_, err = uc.userRepo.Update(ctx, query, data)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Interface("data", data).Msg("unable to update user")
		return errors.Wrapf(err, "unable to update user %d", req.ID)
	}

	return nil
}
