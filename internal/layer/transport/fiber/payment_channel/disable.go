package payment_channel

import (
	"vm-backend/internal/core/infra/network/fiber/http"
	"vm-backend/internal/layer/usecase/payment_channel/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) Disable(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeDisableRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.Disable(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}

func makeDisableRequest(c *fiber.Ctx) (*request.Disable, error) {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	return &request.Disable{ID: uint(id)}, nil
}
