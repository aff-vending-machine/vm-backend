package payment_channel

import (
	"context"

	"vm-backend/internal/layer/usecase/payment_channel/request"
	"vm-backend/internal/layer/usecase/payment_channel/response"
	"vm-backend/pkg/helpers/conv"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) List(ctx context.Context, req *request.Filter) ([]response.Channel, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	entities, err := uc.channelRepo.FindMany(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find many payment channel")
		return nil, errors.Wrap(err, "unable to find many payment channel")
	}

	return conv.ToArray[response.Channel](entities)
}
