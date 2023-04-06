package payment_channel_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Create(ctx context.Context, req *request.Create) (uint, error) {
	if v := validate.Struct(req); !v.Validate() {
		return 0, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	entity := req.ToEntity()
	err := uc.payment_channelRepo.InsertOne(ctx, entity)
	if err != nil {
		return 0, errors.Wrap(err, "unable to insert machine failed")
	}

	return entity.ID, nil
}
