package machine

import (
	"vm-backend/internal/core/infrastructure/network/fiber/http"
	"vm-backend/internal/layer/usecase/machine/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) Count(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeCountRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	res, err := r.usecase.Count(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}

func makeCountRequest(c *fiber.Ctx) (*request.Filter, error) {
	var req request.Filter
	if err := c.QueryParser(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
