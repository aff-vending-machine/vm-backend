package payment_channel_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) List(ctx context.Context, req *request.Filter) ([]response.PaymentChannel, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	entities, err := uc.payment_channelRepo.FindMany(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "unable to find payment channel")
	}

	return response.PaymentChannelEntityToList(entities), nil
}
