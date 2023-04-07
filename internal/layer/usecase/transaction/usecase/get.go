package transaction_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Get(ctx context.Context, req *request.Get) (*response.Transaction, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	entity, err := uc.transactionRepo.FindOne(ctx, filter)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find transaction %d", req.ID)
	}

	return response.ToTransaction(entity), nil
}
