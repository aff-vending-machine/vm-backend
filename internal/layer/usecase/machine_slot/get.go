package machine_slot

import (
	"context"

	"vm-backend/internal/layer/usecase/machine_slot/request"
	"vm-backend/internal/layer/usecase/machine_slot/response"
	"vm-backend/pkg/conv"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Get(ctx context.Context, req *request.Get) (*response.Slot, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	entity, err := uc.machineSlotRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find machine slot")
		return nil, errors.Wrapf(err, "unable to find slot %d in machine %d", req.ID, req.MachineID)
	}

	return conv.StructTo[response.Slot](entity)
}
