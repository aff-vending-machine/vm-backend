package transaction_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/request"
	"github.com/aff-vending-machine/vm-backend/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Count(ctx context.Context, req *request.Filter) (int64, error) {
	ctx, span := trace.Start(ctx)
	defer span.End()

	res, err := w.usecase.Count(ctx, req)
	if err != nil {
		log.Error().Interface("request", req).Err(err).Send()
		trace.RecordError(span, err)
	}

	return res, err
}
