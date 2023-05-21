package payment_channel

import "vm-backend/internal/core/domain/payment"

type transportImpl struct {
	usecase payment.ChannelUsecase
}

func NewTransport(uc payment.ChannelUsecase) *transportImpl {
	return &transportImpl{uc}
}
