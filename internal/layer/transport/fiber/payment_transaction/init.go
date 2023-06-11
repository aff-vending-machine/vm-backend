package payment_transaction

import (
	"vm-backend/internal/core/domain/payment"
)

type transportImpl struct {
	usecase payment.TransactionUsecase
}

func NewTransport(uc payment.TransactionUsecase) *transportImpl {
	return &transportImpl{uc}
}
