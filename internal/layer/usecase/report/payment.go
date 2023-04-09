package report

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Payment(ctx context.Context, req *request.Report) ([]response.Payment, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToPaymentFilter()
	entities, err := uc.transactionRepo.FindMany(ctx, filter)
	if err != nil {
		log.Error().Err(err).Strs("filter", filter).Msg("unable to find transaction")
		return nil, errors.Wrap(err, "unable to find transaction")
	}

	return response.ToPaymentList(entities), nil
}
