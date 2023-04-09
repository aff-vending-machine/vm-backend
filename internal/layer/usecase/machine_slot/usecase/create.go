package machine_slot_usecase

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/model"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/aff-vending-machine/vm-backend/pkg/errs"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Create(ctx context.Context, req *request.Create) (uint, error) {
	if v := validate.Struct(req); !v.Validate() {
		return 0, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	_, err := uc.machineRepo.FindOne(ctx, []string{fmt.Sprintf("id||=||%d", req.MachineID)})
	if err != nil {
		log.Error().Err(err).Msgf("unable to find machine %s", req.MachineID)
		return 0, errors.Wrapf(err, "unable to find machine %s", req.MachineID)
	}

	if req.MachineID == 1 {
		_, err := uc.findProduct(ctx, req)
		if err != nil {
			return 0, err
		}

		entity := req.ToEntity()
		err = uc.machineSlotRepo.InsertOne(ctx, entity)
		if err != nil {
			log.Error().Err(err).Msgf("unable to insert machine slot at machine %s", req.MachineID)
			return 0, errors.Wrapf(err, "unable to insert machine slot at machine %s", req.MachineID)
		}

		return entity.ID, nil
	}

	return 0, fmt.Errorf("limit to single machine")
}

func (uc *usecaseImpl) findProduct(ctx context.Context, req *request.Create) (*model.Product, error) {
	if req.ProductID == 0 {
		return nil, nil
	}

	product, err := uc.productRepo.FindOne(ctx, []string{fmt.Sprintf("id||=||%d", req.ProductID)})
	if errs.Not(err, "record not found") {
		log.Error().Err(err).Msgf("unable to find product %d", req.ProductID)
		return nil, err
	}
	if errs.Is(err, "record not found") {
		return nil, nil
	}

	return &model.Product{
		SKU:      product.SKU,
		Name:     product.Name,
		Type:     product.Type,
		ImageURL: product.ImageURL,
		Price:    product.Price,
	}, nil
}
