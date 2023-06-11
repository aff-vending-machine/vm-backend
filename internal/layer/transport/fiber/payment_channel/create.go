package payment_channel

import (
	"vm-backend/internal/core/infra/network/fiber/http"
	"vm-backend/internal/layer/usecase/payment_channel/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) Create(c *fiber.Ctx) error {
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
	var req request.Create
	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
