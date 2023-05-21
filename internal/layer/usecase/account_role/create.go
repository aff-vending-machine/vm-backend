package account_role

import (
	"context"
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/layer/usecase/account_role/request"

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

	role := makeAccountRole(req)
	_, err := uc.roleRepo.Create(ctx, role)
	if err != nil {
		log.Error().Err(err).Interface("role", role).Msg("unable to insert role")
		return 0, errors.Wrap(err, "unable to insert role")
	}

	return role.ID, nil
}

func makeAccountRole(req *request.Create) *account.Role {
	return &account.Role{
		Name: req.Name,
		// Permissions: req.Permissions,
	}
}
