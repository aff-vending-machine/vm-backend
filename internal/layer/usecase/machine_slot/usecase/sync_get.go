package machine_slot_usecase

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/model"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/aff-vending-machine/vm-backend/pkg/errs"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) SyncGet(ctx context.Context, req *request.Sync) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	machine, err := uc.machineRepo.FindOne(ctx, req.ToMachineFilter())
	if err != nil {
		return errors.Wrapf(err, "unable to find machine %d", req.MachineID)
	}

	slots, err := uc.rpcAPI.GetSlot(ctx, machine.SerialNumber)
	if err != nil {
		return errors.Wrapf(err, "unable to sync real machine %s", machine.SerialNumber)
	}

	mapCase := make(map[string]int)
	mapIndex := make(map[string]int)

	for _, cs := range machine.Slots {
		mapCase[cs.Code] += 1
	}

	for index, ms := range slots {
		mapCase[ms.Code] += 2
		mapIndex[ms.Code] = index
	}

	for code, condition := range mapCase {
		// case #1 only center, remove form machine
		if condition == 1 {
			uc.machineSlotRepo.DeleteMany(ctx, makeCodeFilter(req.MachineID, code))
			continue
		}
		// case #2 only machine, include to machine
		if condition == 2 {
			slot := slots[mapIndex[code]]
			productID := uc.findProductID(ctx, slot)
			uc.machineSlotRepo.InsertOne(ctx, slot.ToEntity(req.MachineID, productID))
			continue
		}

		// case #3 both, adjust by machine
		if condition == 3 {
			slot := slots[mapIndex[code]]
			productID := uc.findProductID(ctx, slot)
			uc.machineSlotRepo.UpdateMany(ctx, makeCodeFilter(req.MachineID, code), slot.ToJson(productID))
			continue
		}
	}

	return nil
}

func (uc *usecaseImpl) findProductID(ctx context.Context, slot model.Slot) uint {
	productID := uint(0)
	if slot.Product != nil {
		product, err := uc.productRepo.FindOne(ctx, []string{fmt.Sprintf("sku:=:%s", slot.Product.SKU)})
		if errs.Is(err, "not found") {
			product = slot.Product.ToEntity()
			err = uc.productRepo.InsertOne(ctx, product)
		}
		if err != nil {
			log.Error().Err(err).Msg("unable to find or create product")
			return 0
		}

		productID = product.ID
	}

	return productID
}
