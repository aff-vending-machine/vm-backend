package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeSync(router fiber.Router, endpoint http.Sync) {
	api := router.Group("sync/:machine_id")

	api.Get("sync/:machine_id", endpoint.GetMachine)
	api.Get("sync/:machine_id/slots", endpoint.GetSlot)
	api.Get("sync/:machine_id/transactions", endpoint.GetTransaction)
	api.Post("sync/:machine_id/slots", endpoint.SetSlot)
}
