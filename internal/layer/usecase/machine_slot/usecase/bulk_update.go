package machine_slot_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) BulkUpdate(ctx context.Context, req *request.BulkUpdate) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	_, err := uc.machineRepo.FindOne(ctx, req.ToMachineFilter())
	if err != nil {
		log.Error().Err(err).Msgf("unable to find machine %s", req.MachineID)
		return errors.Wrapf(err, "unable to find machine %s", req.MachineID)
	}

	for i, data := range req.Data {
		_, err = uc.machineSlotRepo.UpdateMany(ctx, req.ToFilter(data.ID), req.ToJson(i))
		if err != nil {
			log.Error().Err(err).Msgf("unable to update slot %d in machine %d", data.ID, req.MachineID)
			continue
		}
	}

	return nil
}
