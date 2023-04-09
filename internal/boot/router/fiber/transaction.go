package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeTransaction(api fiber.Router, endpoint http.Transaction) {
	api.Get("transactions", endpoint.Read)
	api.Get("transactions/count", endpoint.Count)
	api.Get("transactions/:id", endpoint.ReadOne)
	api.Post("transactions", endpoint.Create)
	api.Put("transactions/:id", endpoint.Update)
	api.Delete("transactions/:id", endpoint.Delete)

	api.Post("transactions/:id/done", endpoint.Done)
	api.Post("transactions/:id/cancel", endpoint.Cancel)
}
