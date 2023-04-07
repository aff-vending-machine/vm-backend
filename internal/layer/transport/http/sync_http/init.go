package sync_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/sync/request"
	"github.com/gofiber/fiber/v2"
)

type httpImpl struct {
	usecase usecase.Sync
}

func New(uc usecase.Sync) *httpImpl {
	return &httpImpl{uc}
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
