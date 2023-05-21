package apiv2

import (
	"vm-backend/internal/core/domain/machine"

	"github.com/gofiber/fiber/v2"
)

func RouteMachine(router fiber.Router, endpoint machine.Transport) {
	api := router.Group("machines")

	api.Get("", endpoint.Read)
	api.Get("count", endpoint.Count)
	api.Get(":id", endpoint.ReadOne)
	api.Post("", endpoint.Create)
	api.Put(":id", endpoint.Update)
	api.Delete(":id", endpoint.Delete)
}

func RouteMachineSlot(router fiber.Router, endpoint machine.SlotTransport) {
	api := router.Group("machines/:machine_id/slots")

	api.Get("", endpoint.Read)
	api.Get("count", endpoint.Count)
	api.Get(":id", endpoint.ReadOne)
	api.Post("", endpoint.Create)
	api.Put("bulk", endpoint.BulkUpdate)
	api.Put(":id", endpoint.Update)
	api.Delete(":id", endpoint.Delete)
}
