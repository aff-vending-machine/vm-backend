package payment_transaction

import (
	"context"

	"vm-backend/internal/layer/usecase/payment_transaction/request"
	"vm-backend/internal/layer/usecase/payment_transaction/response"
	"vm-backend/pkg/conv"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Get(ctx context.Context, req *request.Get) (*response.Transaction, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	entity, err := uc.transactionRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to get transaction")
		return nil, errors.Wrap(err, "unable to get transaction")
	}

	return conv.StructTo[response.Transaction](entity)
}