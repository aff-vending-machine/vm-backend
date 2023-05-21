package payment_transaction

import (
	"context"

	"vm-backend/internal/layer/usecase/payment_transaction/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Done(ctx context.Context, req *request.Done) error {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()

	origin, err := uc.transactionRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find transaction")
		return errors.Wrap(err, "unable to find transaction")
	}

	data := req.ToUpdate(origin.OrderQuantity, origin.OrderPrice, origin.ConfirmedPaidBy)
	_, err = uc.transactionRepo.Update(ctx, query, data)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to update transaction")
		return errors.Wrap(err, "unable to update transaction")
	}

	return nil
}
