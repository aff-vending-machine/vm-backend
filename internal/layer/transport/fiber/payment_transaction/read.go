package payment_transaction

import (
	"vm-backend/internal/core/infrastructure/network/fiber/http"
	"vm-backend/internal/layer/usecase/payment_transaction/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) Read(c *fiber.Ctx) error {
	ctx := c.Context()

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
