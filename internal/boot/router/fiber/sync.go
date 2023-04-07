package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeSync(router fiber.Router, endpoint http.Sync) {
	api := router.Group("sync-machines/:machine_id")

	api.Post("get", endpoint.GetMachine)
	api.Post("slots/get", endpoint.GetSlot)
	api.Post("slots/set", endpoint.SetSlot)
	api.Post("transactions/get", endpoint.GetTransaction)
}
