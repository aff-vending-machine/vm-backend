package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeRole(api fiber.Router, endpoint http.Role) {
	api.Get("roles", endpoint.Read)
	api.Get("roles/count", endpoint.Count)
	api.Get("roles/:id", endpoint.ReadOne)
	api.Post("roles", endpoint.Create)
	api.Put("roles/:id", endpoint.Update)
	api.Delete("roles/:id", endpoint.Delete)
}
