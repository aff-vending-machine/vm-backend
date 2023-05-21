package sync

import (
	"context"
	"time"

	"vm-backend/internal/core/domain/catalog"
	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/core/domain/payment"
	"vm-backend/internal/core/domain/sync"
	"vm-backend/pkg/db"
)

type usecaseImpl struct {
	syncAPI         sync.API
	channelRepo     payment.ChannelRepository
	machineRepo     machine.Repository
	slotRepo        machine.SlotRepository
	productRepo     catalog.ProductRepository
	transactionRepo payment.TransactionRepository
}

func NewUsecase(
	api sync.API,
	pcr payment.ChannelRepository,
	mmr machine.Repository,
	msr machine.SlotRepository,
	cpr catalog.ProductRepository,
	ptr payment.TransactionRepository,
) sync.Usecase {
	return &usecaseImpl{
		api,
		pcr,
		mmr,
		msr,
		cpr,
		ptr,
	}
}

func makeCodeQuery(machineID uint, code string) *db.Query {
	return db.NewQuery().
		AddWhere("machine_id = ?", machineID).
		AddWhere("code = ?", code)
}

func (uc *usecaseImpl) updateMachineStatus(ctx context.Context, query *db.Query, status string) {
	data := map[string]interface{}{
		"status":    status,
		"sync_time": time.Now(),
	}

	uc.machineRepo.Update(ctx, query, data)
}
