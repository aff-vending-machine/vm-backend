package payment_channel_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Get(ctx context.Context, req *request.Get) (*response.PaymentChannel, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	entity, err := uc.payment_channelRepo.FindOne(ctx, filter)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find payment channel %d", req.ID)
	}

	return response.PaymentChannelEntityToView(entity), nil
}
