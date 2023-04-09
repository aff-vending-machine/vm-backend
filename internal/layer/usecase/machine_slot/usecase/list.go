package machine_slot_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) List(ctx context.Context, req *request.Filter) ([]response.MachineSlot, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	entities, err := uc.machineSlotRepo.FindMany(ctx, filter)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find slot in machine %d", req.MachineID)
	}

	return response.ToMachineSlotList(entities), nil
}
