package payment_channel

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel/request"
	"github.com/gofiber/fiber/v2"
)

func (r *restImpl) Active(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeActiveRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.Active(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}

func makeActiveRequest(c *fiber.Ctx) (*request.Active, error) {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	var req request.Active
	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}
	req.ID = uint(id)

	return &req, nil
}
