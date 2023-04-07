package machine_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine/request"
	"github.com/aff-vending-machine/vm-backend/pkg/errs"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) SyncRegister(ctx context.Context, req *request.SyncRegister) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	machine, err := uc.machineRepo.FindOne(ctx, req.ToFilter())
	if errs.Is(err, "not found") {
		err := uc.machineRepo.InsertOne(ctx, req.ToEntity())
		if err != nil {
			return errors.Wrap(err, "unable to insert machine")
		}

		return nil
	}
	if err != nil {
		return errors.Wrapf(err, "unable to find machine (%s)", req.Data.SerialNumber)
	}

	_, err = uc.machineRepo.UpdateMany(ctx, req.ToFilter(), req.ToJsonUpdate(machine.Count))
	if err != nil {
		return errors.Wrapf(err, "unable to update machine (%s)", req.Data.SerialNumber)
	}

	return nil
}
