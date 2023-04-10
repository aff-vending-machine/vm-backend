package product_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/product/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/product/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) List(ctx context.Context, req *request.Filter) ([]response.Product, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()
	entities, err := uc.productRepo.FindMany(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "unable to find product")
	}

	return response.ToProductList(entities), nil
}
