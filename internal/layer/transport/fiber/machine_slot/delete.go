package machine_slot

import (
	"vm-backend/internal/core/infra/network/fiber/http"
	"vm-backend/internal/layer/usecase/machine_slot/request"

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
		return nil, err
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	return &request.Delete{ID: uint(id), MachineID: uint(machineID)}, nil
}
