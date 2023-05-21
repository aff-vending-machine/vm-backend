package payment_transaction

import (
	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/core/domain/payment"
)

type usecaseImpl struct {
	machineRepo     machine.Repository
	transactionRepo payment.TransactionRepository
}

func NewUsecase(
	mmr machine.Repository,
	ptr payment.TransactionRepository,
) *usecaseImpl {
	return &usecaseImpl{
		mmr,
		ptr,
	}
}
