package machine_slot_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Update(ctx context.Context, req *request.Update) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	_, err := uc.machineRepo.FindOne(ctx, req.ToMachineFilter())
	if err != nil {
		log.Error().Err(err).Msgf("unable to find machine %s", req.MachineID)
		return errors.Wrapf(err, "unable to find machine %s", req.MachineID)
	}

	_, err = uc.machineSlotRepo.UpdateMany(ctx, req.ToFilter(), req.ToJson())
	if err != nil {
		log.Error().Err(err).Msgf("unable to update slot %d in machine %d", req.ID, req.MachineID)
		return errors.Wrapf(err, "unable to update slot %d in machine %d", req.ID, req.MachineID)
	}

	return nil
}
