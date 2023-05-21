package payment_channel

import (
	"vm-backend/internal/core/infrastructure/network/fiber/http"
	"vm-backend/internal/layer/usecase/payment_channel/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) Enable(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeEnableRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.Enable(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}

func makeEnableRequest(c *fiber.Ctx) (*request.Enable, error) {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	return &request.Enable{ID: uint(id)}, nil
}
