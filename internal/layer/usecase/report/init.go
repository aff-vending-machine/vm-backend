package report

import (
	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/core/domain/payment"
)

type usecaseImpl struct {
	machineRepo     machine.Repository
	slotRepo        machine.SlotRepository
	channelRepo     payment.ChannelRepository
	transactionRepo payment.TransactionRepository
}

func NewUsecase(
	mmr machine.Repository,
	msr machine.SlotRepository,
	pcr payment.ChannelRepository,
	ptr payment.TransactionRepository,
) *usecaseImpl {
	return &usecaseImpl{
		mmr,
		msr,
		pcr,
		ptr,
	}
}
