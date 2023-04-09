package report

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Transaction(ctx context.Context, req *request.Report) ([]response.Transaction, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToTransactionFilter()
	entities, err := uc.transactionRepo.FindMany(ctx, filter)
	if err != nil {
		log.Error().Err(err).Strs("filter", filter).Msg("unable to find transaction")
		return nil, errors.Wrap(err, "unable to find transaction")
	}

	return response.ToTransactionList(entities), nil
}
