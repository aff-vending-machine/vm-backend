package product_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/product/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/product/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Get(ctx context.Context, req *request.Get) (*response.Product, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	entity, err := uc.productRepo.FindOne(ctx, filter)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find product %d", req.ID)
	}

	return response.ToProduct(entity), nil
}
