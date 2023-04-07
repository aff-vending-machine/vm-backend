package transaction_usecase

import "github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"

type usecaseImpl struct {
	transactionRepo repository.Transaction
}

func New(p repository.Transaction) *usecaseImpl {
	return &usecaseImpl{p}
}
