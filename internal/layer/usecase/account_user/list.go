package account_user

import (
	"context"

	"vm-backend/internal/layer/usecase/account_user/request"
	"vm-backend/internal/layer/usecase/account_user/response"
	"vm-backend/pkg/conv"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) List(ctx context.Context, req *request.Filter) ([]response.User, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	entities, err := uc.userRepo.FindMany(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find users")
		return nil, errors.Wrap(err, "unable to find users")
	}

	return conv.StructToArray[response.User](entities)
}
