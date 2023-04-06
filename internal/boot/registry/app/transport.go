package app

import (
	"github.com/aff-vending-machine/vm-backend/internal/boot/registry"
)

func NewTransport(uc registry.Usecase) registry.Transport {
	return registry.Transport{}
}
