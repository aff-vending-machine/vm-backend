package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeAuth(api fiber.Router, endpoint http.Auth) {
	api.Post("auth/login", endpoint.Login)
	api.Get("auth/token", endpoint.RefreshToken)
}
