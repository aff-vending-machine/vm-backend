package machine_slot

import (
	"context"

	"vm-backend/internal/layer/usecase/machine_slot/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Update(ctx context.Context, req *request.Update) error {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return errors.Wrap(err, "unable to validate request")
	}

	if isExist := uc.isMachineExist(ctx, req.MachineID); !isExist {
		return errors.Errorf("machine %d not found", req.MachineID)
	}

	query := req.ToQuery()
	update := req.ToUpdate()
	_, err := uc.machineSlotRepo.Update(ctx, query, update)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Interface("data", update).Msg("unable to update machine slot")
		return errors.Wrapf(err, "unable to update machine slot %d", req.ID)
	}

	return nil
}
