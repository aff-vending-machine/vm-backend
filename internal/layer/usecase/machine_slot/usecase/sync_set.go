package machine_slot_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/model"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) SyncSet(ctx context.Context, req *request.Sync) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	machine, err := uc.machineRepo.FindOne(ctx, req.ToMachineFilter())
	if err != nil {
		return errors.Wrapf(err, "unable to find machine %s", req.MachineID)
	}

	slots, err := uc.machineSlotRepo.FindMany(ctx, req.ToFilter())
	if err != nil {
		return errors.Wrapf(err, "unable to find slot in machine %d", req.MachineID)
	}

	err = uc.rpcAPI.SlotSet(ctx, machine.SerialNumber, model.ToSlotList(slots))
	if err != nil {
		return errors.Wrapf(err, "unable to sync set to real machine %s", machine.SerialNumber)
	}

	return nil
}
