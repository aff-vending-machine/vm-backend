package machine_slot

import (
	"context"
	"vm-backend/internal/layer/usecase/machine_slot/request"

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
	total, err := uc.machineSlotRepo.Count(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to count machine slots")
		return 0, errors.Wrap(err, "unable to count machine slots")
	}

	return total, nil
}
