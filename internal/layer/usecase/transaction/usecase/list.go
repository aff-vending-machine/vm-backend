package transaction_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) List(ctx context.Context, req *request.Filter) ([]response.Transaction, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	entities, err := uc.transactionRepo.FindMany(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "unable to find transaction")
	}

	return response.ToTransactionList(entities), nil
}
