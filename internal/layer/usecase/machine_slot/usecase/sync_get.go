package machine_slot_usecase

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) SyncGet(ctx context.Context, req *request.Sync) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	machine, err := uc.machineRepo.FindOne(ctx, req.ToMachineFilter())
	if err != nil {
		return errors.Wrapf(err, "unable to find machine %s", req.MachineID)
	}

	slots, err := uc.rpcAPI.SlotGet(ctx, machine.SerialNumber)
	if err != nil {
		return errors.Wrapf(err, "unable to sync real machine %s", machine.SerialNumber)
	}

	mapCase := make(map[string]int, 1)
	mapIndex := make(map[string]int, 1)

	for index, cs := range machine.Slots {
		mapCase[cs.Code] += 1
		mapIndex[cs.Code] = index
	}

	for index, ms := range slots {
		mapCase[ms.Code] += 2
		mapIndex[ms.Code] = index
	}

	for code, condition := range mapCase {
		// case #1 only center, remove form machine
		if condition == 1 {
			uc.machineSlotRepo.DeleteMany(ctx, makeCodeFilter(req.MachineID, code))
		}
		// case #2 only machine, include to machine
		if condition == 2 {
			productID := uint(0)
			slot := slots[mapIndex[code]]
			if slot.Product != nil {
				product, _ := uc.productRepo.FindOne(ctx, []string{fmt.Sprintf("sku:=:%s", slot.Product.SKU)})
				productID = product.ID
			}
			uc.machineSlotRepo.InsertOne(ctx, slot.ToEntity(req.MachineID, productID))
		}

		// case #3 both, adjust by machine
		if condition == 3 {
			productID := uint(0)
			slot := slots[mapIndex[code]]
			if slot.Product != nil {
				product, _ := uc.productRepo.FindOne(ctx, []string{fmt.Sprintf("sku:=:%s", slot.Product.SKU)})
				productID = product.ID
			}
			uc.machineSlotRepo.UpdateMany(ctx, makeCodeFilter(req.MachineID, code), slot.ToJson(productID))
		}
	}

	return nil
}
