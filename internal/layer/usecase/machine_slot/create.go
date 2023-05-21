package machine_slot

import (
	"context"

	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/layer/usecase/machine_slot/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Create(ctx context.Context, req *request.Create) (uint, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return 0, errors.Wrap(err, "unable to validate request")
	}

	if isExist := uc.isMachineExist(ctx, req.MachineID); !isExist {
		return 0, errors.Errorf("machine %s not found", req.MachineID)
	}

	if req.ProductID != 0 {
		if isExist := uc.isProductExist(ctx, req.ProductID); !isExist {
			return 0, errors.Errorf("product %s not found", req.ProductID)
		}
	}

	entity := makeMachineSlot(req)
	_, err := uc.machineSlotRepo.Create(ctx, entity)
	if err != nil {
		log.Error().Err(err).Interface("entity", entity).Msg("unable to create machine slot")
		return 0, errors.Wrapf(err, "unable to create machine slot at machine %s", req.MachineID)
	}

	return entity.ID, nil
}

func makeMachineSlot(req *request.Create) *machine.Slot {
	return &machine.Slot{
		MachineID:        req.MachineID,
		CatalogProductID: req.ProductID,
		Code:             req.Code,
		Stock:            req.Stock,
		Capacity:         req.Capacity,
		IsEnable:         req.IsEnable,
	}
}
