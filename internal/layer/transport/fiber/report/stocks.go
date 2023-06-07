package report

import (
	"vm-backend/internal/core/infra/network/fiber/http"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) Stocks(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeReportRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	res, err := r.usecase.Stocks(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}
