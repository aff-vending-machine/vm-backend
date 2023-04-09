package machine_slot

import (
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/request"
	"github.com/gofiber/fiber/v2"
)

func (r *restImpl) Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeDeleteRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.Delete(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}

func makeDeleteRequest(c *fiber.Ctx) (*request.Delete, error) {
	machineID, err := c.ParamsInt("machine_id", 0)
	if err != nil {
		return nil, fmt.Errorf("invalid Machine ID")
	}
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	return &request.Delete{MachineID: uint(machineID), ID: uint(id)}, nil
}
