package transaction_wrapper

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction"
)

type wrapperImpl struct {
	usecase transaction.Usecase
}

func New(usecase transaction.Usecase) transaction.Usecase {
	return &wrapperImpl{usecase}
}
