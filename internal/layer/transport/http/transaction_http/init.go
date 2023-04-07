package transaction_http

import "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction"

type restImpl struct {
	usecase transaction.Usecase
}

func New(uc transaction.Usecase) *restImpl {
	return &restImpl{uc}
}
