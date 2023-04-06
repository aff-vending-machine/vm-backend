package machine_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) List(ctx context.Context, req *request.Filter) ([]response.Machine, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	entities, err := uc.machineRepo.FindMany(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "unable to find machines")
	}

	return response.MachineEntityToList(entities), nil
}
