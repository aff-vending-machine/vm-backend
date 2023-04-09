package transaction_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase"
)

type httpImpl struct {
	usecase usecase.Transaction
}

func New(uc usecase.Transaction) *httpImpl {
	return &httpImpl{uc}
}
