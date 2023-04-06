package machine_slot_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Delete(ctx context.Context, req *request.Delete) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToMachineFilter()
	_, err := uc.machineRepo.FindOne(ctx, filter)
	if err != nil {
		return errors.Wrapf(err, "unable to find machine %d", req.MachineID)
	}

	filter = req.ToFilter()
	_, err = uc.machineSlotRepo.FindOne(ctx, filter)
	if err != nil {
		return errors.Wrapf(err, "unable to find slot %d at machine %d", req.ID, req.MachineID)
	}

	filter = req.ToFilter()
	_, err = uc.machineRepo.DeleteMany(ctx, filter)
	if err != nil {
		return errors.Wrapf(err, "unable to delete slot %d in machine %d", req.ID, req.MachineID)
	}

	return nil
}
