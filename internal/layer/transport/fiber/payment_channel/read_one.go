package payment_channel

import (
	"vm-backend/internal/core/infra/network/fiber/http"
	"vm-backend/internal/layer/usecase/payment_channel/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) ReadOne(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeGetRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	res, err := r.usecase.Get(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}

func makeGetRequest(c *fiber.Ctx) (*request.Get, error) {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	return &request.Get{ID: uint(id)}, nil
}
