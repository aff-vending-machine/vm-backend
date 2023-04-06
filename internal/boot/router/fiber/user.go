package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeUser(api fiber.Router, endpoint http.User) {
	api.Get("users", endpoint.Read)
	api.Get("users/count", endpoint.Count)
	api.Get("users/:id", endpoint.ReadOne)
	api.Post("users", endpoint.Create)
	api.Post("users/:id/change-role", endpoint.ChangeRole)
	api.Post("users/me/change-password", endpoint.ChangePassword)
	api.Post("users/:id/reset-password", endpoint.ResetPassword)
	// api.Put("users", endpoint.Update)
	api.Delete("users/:id", endpoint.Delete)
}
