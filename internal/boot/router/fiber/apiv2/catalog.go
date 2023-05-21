package apiv2

import (
	"vm-backend/internal/core/domain/catalog"

	"github.com/gofiber/fiber/v2"
)

func RouteCatalogGroup(router fiber.Router, endpoint catalog.GroupTransport) {
	api := router.Group("groups")

	api.Get("", endpoint.Read)
	api.Get("count", endpoint.Count)
	api.Get(":id", endpoint.ReadOne)
	api.Post("", endpoint.Create)
	api.Put(":id", endpoint.Update)
	api.Delete(":id", endpoint.Delete)
}

func RouteCatalogProduct(router fiber.Router, endpoint catalog.ProductTransport) {
	api := router.Group("products")

	api.Get("", endpoint.Read)
	api.Get("count", endpoint.Count)
	api.Get(":id", endpoint.ReadOne)
	api.Post("", endpoint.Create)
	api.Put(":id", endpoint.Update)
	api.Delete(":id", endpoint.Delete)
}
