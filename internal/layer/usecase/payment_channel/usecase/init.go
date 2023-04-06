package payment_channel_usecase

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

type usecaseImpl struct {
	payment_channelRepo repository.PaymentChannel
}

func New(p repository.PaymentChannel) *usecaseImpl {
	return &usecaseImpl{p}
}
