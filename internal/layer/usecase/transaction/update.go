package transaction

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Update(ctx context.Context, req *request.Update) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	data := req.ToJson()
	_, err := uc.transactionRepo.UpdateMany(ctx, filter, data)
	if err != nil {
		return errors.Wrap(err, "unable to update transaction")
	}

	return nil
}
