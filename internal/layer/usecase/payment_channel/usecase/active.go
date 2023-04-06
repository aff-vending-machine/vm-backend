package payment_channel_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Active(ctx context.Context, req *request.Active) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	data := req.ToJson()
	_, err := uc.payment_channelRepo.UpdateMany(ctx, filter, data)
	if err != nil {
		return errors.Wrapf(err, "unable to update payment channel %d", req.ID)
	}

	return nil
}
