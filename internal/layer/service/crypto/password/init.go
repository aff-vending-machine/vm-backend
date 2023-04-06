package password

import (
	"github.com/aff-vending-machine/vm-backend/config"
)

type managerImpl struct {
	salt int
}

func New(config config.BCryptConfig) *managerImpl {
	return &managerImpl{config.Salt}
}
