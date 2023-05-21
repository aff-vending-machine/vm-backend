package payment_channel

import (
	"vm-backend/internal/core/infrastructure/network/fiber/http"
	"vm-backend/internal/layer/usecase/payment_channel/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) Update(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeUpdateRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.Update(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}

func makeUpdateRequest(c *fiber.Ctx) (*request.Update, error) {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	var req request.Update
	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}
	req.ID = uint(id)

	return &req, nil
}
