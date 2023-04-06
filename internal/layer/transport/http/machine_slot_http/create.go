package machine_slot_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/gofiber/fiber/v2"
)

func (r *restImpl) Create(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeCreateRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	id, err := r.usecase.Create(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.Created(c, id)
}

func makeCreateRequest(c *fiber.Ctx) (*request.Create, error) {
	machineID, err := c.ParamsInt("machine_id", 0)
	if err != nil {
		return nil, err
	}

	var req request.Create
	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}
	req.MachineID = uint(machineID)

	return &req, nil
}
