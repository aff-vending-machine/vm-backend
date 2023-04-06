package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routePaymentChannel(api fiber.Router, endpoint http.PaymentChannel) {
	api.Get("payment_channels", endpoint.Read)
	api.Get("payment_channels/count", endpoint.Count)
	api.Get("payment_channels/:id", endpoint.ReadOne)
	api.Post("payment_channels/:id/active", endpoint.Active)
	api.Put("payment_channels/:id", endpoint.Update)
	api.Delete("payment_channels/:id", endpoint.Delete)
}
