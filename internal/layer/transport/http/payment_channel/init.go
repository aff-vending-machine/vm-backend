package payment_channel

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel"
)

type restImpl struct {
	usecase payment_channel.Usecase
}

func New(uc payment_channel.Usecase) *restImpl {
	return &restImpl{uc}
}
