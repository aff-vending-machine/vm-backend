package apiv2

import (
	"vm-backend/internal/core/domain/store"

	"github.com/gofiber/fiber/v2"
)

func RouteStoreBranch(router fiber.Router, endpoint store.BranchTransport) {
	api := router.Group("branches")

	api.Get("", endpoint.Read)
	api.Get("count", endpoint.Count)
	api.Get(":id", endpoint.ReadOne)
	api.Post("", endpoint.Create)
	api.Put(":id", endpoint.Update)
	api.Delete(":id", endpoint.Delete)
}
