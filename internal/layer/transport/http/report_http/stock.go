package report_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/gofiber/fiber/v2"
)

func (r *httpImpl) GetStock(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeReportRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	res, err := r.usecase.Stock(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}
