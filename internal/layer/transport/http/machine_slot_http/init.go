package machine_slot_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/gofiber/fiber/v2"
)

type restImpl struct {
	usecase machine_slot.Usecase
}

func New(uc machine_slot.Usecase) *restImpl {
	return &restImpl{uc}
}

func makeSyncRequest(c *fiber.Ctx) (*request.Sync, error) {
	machineID, err := c.ParamsInt("machine_id", 0)
	if err != nil {
		return nil, err
	}

	req := request.Sync{
		MachineID: uint(machineID),
	}

	return &req, nil
}
