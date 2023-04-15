package machine_slot

import (
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/gofiber/fiber/v2"
)

func (r *restImpl) BulkUpdate(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeBulkUpdateRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.BulkUpdate(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}

func makeBulkUpdateRequest(c *fiber.Ctx) (*request.BulkUpdate, error) {
	machineID, err := c.ParamsInt("machine_id", 0)
	if err != nil {
		return nil, fmt.Errorf("invalid machine ID")
	}

	var req request.BulkUpdate
	if err := c.BodyParser(&req.Data); err != nil {
		return nil, err
	}
	req.MachineID = uint(machineID)

	return &req, nil
}
