package machine

import "vm-backend/internal/core/domain/machine"

type transportImpl struct {
	usecase machine.Usecase
}

func NewTransport(uc machine.Usecase) machine.Transport {
	return &transportImpl{uc}
}
