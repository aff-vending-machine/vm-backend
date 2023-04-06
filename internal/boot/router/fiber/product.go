package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeProduct(api fiber.Router, endpoint http.Product) {
	api.Get("products", endpoint.Read)
	api.Get("products/count", endpoint.Count)
	api.Get("products/:id", endpoint.ReadOne)
	api.Post("products", endpoint.Create)
	api.Put("products/:id", endpoint.Update)
	api.Delete("products/:id", endpoint.Delete)
}
