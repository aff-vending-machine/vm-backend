package transaction_usecase

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

type usecaseImpl struct {
	rpcAPI          api.RPC
	machineRepo     repository.Machine
	transactionRepo repository.Transaction
}

func New(r api.RPC, m repository.Machine, t repository.Transaction) *usecaseImpl {
	return &usecaseImpl{r, m, t}
}
