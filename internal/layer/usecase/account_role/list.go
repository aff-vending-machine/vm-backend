package account_role

import (
	"context"

	"vm-backend/internal/layer/usecase/account_role/request"
	"vm-backend/internal/layer/usecase/account_role/response"
	"vm-backend/pkg/conv"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) List(ctx context.Context, req *request.Filter) ([]response.Role, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	entities, err := uc.roleRepo.FindMany(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find roles")
		return nil, errors.Wrap(err, "unable to find roles")
	}

	return conv.StructToArray[response.Role](entities)
}