package payment_channel

import (
	"vm-backend/internal/core/domain/payment"
)

type usecaseImpl struct {
	channelRepo payment.ChannelRepository
}

func NewUsecase(p payment.ChannelRepository) payment.ChannelUsecase {
	return &usecaseImpl{p}
}
