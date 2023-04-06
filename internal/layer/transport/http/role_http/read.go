package role_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role/request"
	"github.com/gofiber/fiber/v2"
)

func (r *restImpl) Read(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeListRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	list, err := r.usecase.List(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.OK(c, list)
}

func makeListRequest(c *fiber.Ctx) (*request.Filter, error) {
	var req request.Filter
	if err := c.QueryParser(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
