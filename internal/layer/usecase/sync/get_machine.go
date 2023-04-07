package sync

import (
	"context"
	"strings"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/sync/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) GetMachine(ctx context.Context, req *request.Sync) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToMachineFilter()
	machine, err := uc.machineRepo.FindOne(ctx, filter)
	if err != nil {
		return errors.Wrapf(err, "unable to find machine %s", req.MachineID)
	}

	res, err := uc.rpcAPI.GetMachine(ctx, machine.SerialNumber)
	if err != nil && strings.Contains(err.Error(), "internal") {
		machine.Status = "maintenance"
		uc.updateMachineStatus(ctx, filter, "maintenance")
		return errors.Wrapf(err, "real machine %s is not ready", machine.SerialNumber)
	}
	if err != nil {
		machine.Status = "offline"
		uc.updateMachineStatus(ctx, filter, "offline")
		return errors.Wrapf(err, "unable to sync real machine %s", machine.SerialNumber)
	}

	uc.machineRepo.UpdateMany(ctx, filter, res.ToJsonUpdate())

	return nil
}
