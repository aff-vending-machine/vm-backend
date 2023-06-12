package machine_slot

import (
	"context"

	"vm-backend/internal/layer/usecase/machine_slot/request"
	"vm-backend/internal/layer/usecase/machine_slot/response"
	"vm-backend/pkg/helpers/conv"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) List(ctx context.Context, req *request.Filter) ([]response.Slot, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	entities, err := uc.machineSlotRepo.FindMany(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find many machine slot")
		return nil, errors.Wrap(err, "unable to find many machine slot")
	}

	return conv.ToArray[response.Slot](entities)
}
