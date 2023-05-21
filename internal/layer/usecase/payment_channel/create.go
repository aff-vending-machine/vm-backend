package payment_channel

import (
	"context"

	"vm-backend/internal/core/domain/payment"
	"vm-backend/internal/layer/usecase/payment_channel/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Create(ctx context.Context, req *request.Create) (uint, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return 0, errors.Wrap(err, "unable to validate request")
	}

	entity := makeChannel(req)
	_, err := uc.channelRepo.Create(ctx, entity)
	if err != nil {
		log.Error().Err(err).Interface("entity", entity).Msg("unable to create payment channel")
		return 0, errors.Wrap(err, "unable to create payment channel")
	}

	return entity.ID, nil
}

func makeChannel(req *request.Create) *payment.Channel {
	return &payment.Channel{
		Name:         req.Name,
		Channel:      req.Channel,
		Vendor:       req.Vendor,
		IsEnable:     req.IsEnable,
		Host:         req.Host,
		MerchantID:   req.MerchantID,
		MerchantName: req.MerchantName,
		BillerCode:   req.BillerCode,
		BillerID:     req.BillerID,
		StoreID:      req.StoreID,
		TerminalID:   req.TerminalID,
	}
}
