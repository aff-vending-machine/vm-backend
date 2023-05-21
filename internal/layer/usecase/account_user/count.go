package account_user

import (
	"context"
	"vm-backend/internal/layer/usecase/account_user/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Count(ctx context.Context, req *request.Filter) (int64, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return 0, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	total, err := uc.userRepo.Count(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to count user")
		return 0, errors.Wrap(err, "unable to count user")
	}

	return total, nil
}
