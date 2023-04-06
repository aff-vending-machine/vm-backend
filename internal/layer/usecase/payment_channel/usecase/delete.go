package payment_channel_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Delete(ctx context.Context, req *request.Delete) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	_, err := uc.payment_channelRepo.DeleteMany(ctx, filter)
	if err != nil {
		return errors.Wrapf(err, "unable to delete machine %d", req.ID)
	}

	return nil
}
