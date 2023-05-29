package machine_slot

import (
	"context"
	"vm-backend/internal/core/domain/catalog"
	"vm-backend/internal/core/domain/machine"
	"vm-backend/pkg/db"

	"github.com/rs/zerolog/log"
)

type usecaseImpl struct {
	machineRepo        machine.Repository
	machineSlotRepo    machine.SlotRepository
	catalogProductRepo catalog.ProductRepository
}

func NewUsecase(
	mmr machine.Repository,
	msr machine.SlotRepository,
	cpr catalog.ProductRepository,
) machine.SlotUsecase {
	return &usecaseImpl{
		mmr,
		msr,
		cpr,
	}
}

func (uc *usecaseImpl) isMachineExist(ctx context.Context, id uint) bool {
	query := db.NewQuery().AddWhere("id = ?", id)
	_, err := uc.machineRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find machine")
		return false
	}

	return true
}

func (uc *usecaseImpl) isProductExist(ctx context.Context, id uint) bool {
	query := db.NewQuery().AddWhere("id = ?", id)
	_, err := uc.catalogProductRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find catalog product")
		return false
	}

	return true
}
