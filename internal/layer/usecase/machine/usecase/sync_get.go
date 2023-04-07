package machine_usecase

import (
	"context"
	"strings"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) SyncGet(ctx context.Context, req *request.Sync) (*response.MachineStatus, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	machine, err := uc.machineRepo.FindOne(ctx, filter)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find machine %s", req.ID)
	}

	res, err := uc.rpcAPI.GetMachine(ctx, machine.SerialNumber)
	if err != nil && strings.Contains(err.Error(), "internal") {
		machine.Status = "maintenance"
		uc.updateMachineStatus(ctx, filter, "maintenance")
		return nil, errors.Wrapf(err, "real machine %s is not ready", machine.SerialNumber)
	}
	if err != nil {
		machine.Status = "offline"
		uc.updateMachineStatus(ctx, filter, "offline")
		return nil, errors.Wrapf(err, "unable to sync real machine %s", machine.SerialNumber)
	}

	uc.machineRepo.UpdateMany(ctx, filter, res.ToJsonUpdate())

	return response.ToMachineStatus(machine), nil
}
