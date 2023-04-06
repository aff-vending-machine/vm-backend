package machine_slot_usecase

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Update(ctx context.Context, req *request.Update) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	_, err := uc.machineRepo.FindOne(ctx, []string{fmt.Sprintf("id:=:%d", req.MachineID)})
	if err != nil {
		log.Error().Err(err).Msgf("unable to find machine %s", req.MachineID)
		return errors.Wrapf(err, "unable to find machine %s", req.MachineID)
	}

	_, err = uc.machineSlotRepo.FindOne(ctx, []string{fmt.Sprintf("machine_id:=:%d", req.MachineID), fmt.Sprintf("id:=:%d", req.ID)})
	if err != nil {
		log.Error().Err(err).Msgf("unable to find slot %d in machine %d", req.ID, req.MachineID)
		return errors.Wrapf(err, "unable to find slot %d in machine %d", req.ID, req.MachineID)
	}

	if req.MachineID == 1 {
		updated := map[string]interface{}{}

		if req.ProductID != nil {
			product := &entity.Product{}
			if *req.ProductID == 0 {
				product = &entity.Product{
					SKU:      "",
					Name:     "",
					ImageURL: "",
					Type:     "",
					Price:    0,
				}
			} else {
				p, err := uc.productRepo.FindOne(ctx, []string{fmt.Sprintf("id:=:%d", req.ProductID)})
				if err != nil {
					log.Error().Err(err).Msgf("unable to find product %d in machine slot %d", req.ProductID, req.ID)
					return errors.Wrapf(err, "unable to find product %d in machine slot %d", req.ProductID, req.ID)
				}
				product = p
			}

			updated["product_sku"] = product.SKU
			updated["product_name"] = product.Name
			updated["product_type"] = product.Type
			updated["product_image_url"] = product.ImageURL
			updated["product_price"] = product.Price
		}
		if req.Stock != nil {
			updated["stock"] = req.Stock
		}
		if req.Capacity != nil {
			updated["capacity"] = req.Capacity
		}
		if req.IsEnable != nil {
			updated["is_enable"] = req.IsEnable
		}

		_, err = uc.machineSlotRepo.UpdateMany(ctx, req.ToFilter(), req.ToJson())
		if err != nil {
			log.Error().Err(err).Msgf("unable to update slot %d in machine %d", req.ID, req.MachineID)
			return errors.Wrapf(err, "unable to update slot %d in machine %d", req.ID, req.MachineID)
		}
	}

	return nil
}
