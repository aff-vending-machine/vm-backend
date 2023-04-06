package machine_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Count(ctx context.Context, req *request.Filter) (int64, error) {
	if v := validate.Struct(req); !v.Validate() {
		return 0, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	total, err := uc.machineRepo.Count(ctx, filter)
	if err != nil {
		return 0, errors.Wrap(err, "unable to count machines")
	}

	return total, nil
}
