package machine

import "vm-backend/internal/core/domain/machine"

type usecaseImpl struct {
	machineRepo machine.Repository
}

func NewUsecase(
	mmr machine.Repository,
) machine.Usecase {
	return &usecaseImpl{
		mmr,
	}
}
