package sync

import (
	"context"
	"time"

	"vm-backend/internal/core/domain/sync/models"
	"vm-backend/internal/layer/usecase/sync/request"
	"vm-backend/pkg/db"
	"vm-backend/pkg/errs"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) FetchSlots(ctx context.Context, req *request.Sync) error {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return errors.Wrap(err, "unable to validate request")
	}

	query := req.ToMachineQuery()
	machine, err := uc.machineRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find machine")
		return errors.Wrap(err, "unable to find machine")
	}

	slots, err := uc.syncAPI.GetSlots(ctx, machine.SerialNumber)
	if err != nil {
		log.Error().Err(err).Str("target", machine.SerialNumber).Msg("unable to fetch slots")
		return errors.Wrap(err, "unable to fetch slots")
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
			uc.slotRepo.Delete(ctx, makeCodeQuery(req.MachineID, code))
			continue
		}
		// case #2 only machine, include to machine
		if condition == 2 {
			slot := slots[mapIndex[code]]
			productID := uc.findProductID(ctx, slot)
			uc.slotRepo.Create(ctx, slot.ToDomain(req.MachineID, productID))
			continue
		}

		// case #3 both, adjust by machine
		if condition == 3 {
			slot := slots[mapIndex[code]]
			productID := uc.findProductID(ctx, slot)
			uc.slotRepo.Update(ctx, makeCodeQuery(req.MachineID, code), slot.ToUpdate(productID))
			continue
		}
	}

	query = req.ToMachineQuery()
	update := map[string]interface{}{"sync_slot_time": time.Now()}
	_, err = uc.machineRepo.Update(ctx, query, update)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Interface("update", update).Msg("unable to update machine")
		return errors.Wrap(err, "unable to update machine")
	}

	return nil
}

func (uc *usecaseImpl) findProductID(ctx context.Context, slot models.Slot) uint {
	productID := uint(0)
	if slot.Product != nil {
		query := db.NewQuery().AddWhere("sku = ?", slot.Product.SKU)
		product, err := uc.productRepo.FindOne(ctx, query)
		if errs.Is(err, errs.ErrNotFound) {
			product = slot.Product.ToEntity()
			_, err = uc.productRepo.Create(ctx, product)
		}
		if err != nil {
			log.Error().Err(err).Msg("unable to find or create product")
			return 0
		}

		productID = product.ID
	}

	return productID
}
