package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeMachineSlot(router fiber.Router, endpoint http.MachineSlot) {
	api := router.Group("machines/:machine_id")

	api.Get("slots", endpoint.Read)
	api.Get("slots/count", endpoint.Count)
	api.Get("slots/:id", endpoint.ReadOne)
	api.Post("slots", endpoint.Create)
	api.Put("slots/bulk", endpoint.BulkUpdate)
	api.Put("slots/:id", endpoint.Update)
	api.Delete("slots/:id", endpoint.Delete)
}
