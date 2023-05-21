package apiv2

import (
	"vm-backend/internal/core/domain/report"

	"github.com/gofiber/fiber/v2"
)

func RouteReport(router fiber.Router, endpoint report.Transport) {
	api := router.Group("report")

	api.Get("summary", endpoint.Summary)
	api.Get(":machine_id/transactions", endpoint.Transactions)
	api.Get(":machine_id/stocks", endpoint.Stocks)
}
