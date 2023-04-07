package transaction_http

import "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction"

type httpImpl struct {
	usecase transaction.Usecase
}

func New(uc transaction.Usecase) *httpImpl {
	return &httpImpl{uc}
}
