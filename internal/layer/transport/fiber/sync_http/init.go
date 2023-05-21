package sync_http

import (
	"vm-backend/internal/core/domain/sync"
	"vm-backend/internal/layer/usecase/sync/request"

	"github.com/gofiber/fiber/v2"
)

type transportImpl struct {
	usecase sync.Usecase
}

func NewTransport(uc sync.Usecase) sync.HTTPTransport {
	return &transportImpl{uc}
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
