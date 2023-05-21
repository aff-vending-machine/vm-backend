package apiv2

import (
	"vm-backend/internal/core/domain/account"

	"github.com/gofiber/fiber/v2"
)

func RouteAccount(router fiber.Router, endpoint account.Transport) {
	api := router.Group("auth")

	api.Post("login", endpoint.Login)
	api.Post("refresh-token", endpoint.RefreshToken)
}

func RouteAccountRole(router fiber.Router, endpoint account.RoleTransport) {
	api := router.Group("roles")

	api.Get("", endpoint.Read)
	api.Get("count", endpoint.Count)
	api.Get(":id", endpoint.ReadOne)
	api.Post("", endpoint.Create)
	api.Put(":id", endpoint.Update)
	api.Delete(":id", endpoint.Delete)
}

func RouteAccountUser(router fiber.Router, endpoint account.UserTransport) {
	api := router.Group("users")

	api.Get("", endpoint.Read)
	api.Get("count", endpoint.Count)
	api.Get(":id", endpoint.ReadOne)
	api.Post("", endpoint.Create)
	api.Post(":id/change-role", endpoint.ChangeRole)
	api.Post("me/change-password", endpoint.ChangePassword)
	api.Post(":id/reset-password", endpoint.ResetPassword)
	api.Delete(":id", endpoint.Delete)
}
