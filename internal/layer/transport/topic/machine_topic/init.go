package machine_topic

import "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine"

type machineImpl struct {
	usecase machine.Usecase
}

func New(uc machine.Usecase) *machineImpl {
	return &machineImpl{uc}
}
