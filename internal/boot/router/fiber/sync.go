package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeSync(router fiber.Router, endpoint http.Sync) {
	router.Get("sync/:machine_id", endpoint.GetMachine)
	router.Get("sync/:machine_id/slots", endpoint.GetSlot)
	router.Get("sync/:machine_id/transactions", endpoint.GetTransaction)
	router.Post("sync/:machine_id/slots", endpoint.SetSlot)
}
