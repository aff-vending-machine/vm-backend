package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeMachine(api fiber.Router, endpoint http.Machine) {
	api.Get("machines", endpoint.Read)
	api.Get("machines/count", endpoint.Count)
	api.Get("machines/:id", endpoint.ReadOne)
	api.Post("machines", endpoint.Create)
	api.Put("machines/:id", endpoint.Update)
	api.Delete("machines/:id", endpoint.Delete)
}
