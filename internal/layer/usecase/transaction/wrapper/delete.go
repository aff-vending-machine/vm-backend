package transaction_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/request"
	"github.com/aff-vending-machine/vm-backend/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Delete(ctx context.Context, req *request.Delete) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	err := w.usecase.Delete(ctx, req)
	if err != nil {
		log.Error().Interface("request", req).Err(err).Send()
		trace.RecordError(span, err)
	}

	return err
}
