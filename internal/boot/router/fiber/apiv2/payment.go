package apiv2

import (
	"vm-backend/internal/core/domain/payment"

	"github.com/gofiber/fiber/v2"
)

func RoutePaymentChannel(router fiber.Router, endpoint payment.ChannelTransport) {
	api := router.Group("channels")

	api.Get("", endpoint.Read)
	api.Get("count", endpoint.Count)
	api.Get(":id", endpoint.ReadOne)
	api.Post("", endpoint.Create)
	api.Post(":id/enable", endpoint.Enable)
	api.Post(":id/disable", endpoint.Disable)
	api.Put(":id", endpoint.Update)
	api.Delete(":id", endpoint.Delete)
}

func RoutePaymentTransaction(router fiber.Router, endpoint payment.TransactionTransport) {
	api := router.Group("transactions")

	api.Get("", endpoint.Read)
	api.Get("count", endpoint.Count)
	api.Get(":id", endpoint.ReadOne)
	api.Post(":id/cancel", endpoint.Cancel)
	api.Post(":id/done", endpoint.Done)
}
