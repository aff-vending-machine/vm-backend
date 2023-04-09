package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeReport(api fiber.Router, endpoint http.Report) {
	api.Get("report/payment", endpoint.GetPayment)
	api.Get("report/payment/download", endpoint.DownloadPayment)
	api.Get("report/stock", endpoint.GetStock)
	api.Get("report/stock/download", endpoint.DownloadStock)
}
