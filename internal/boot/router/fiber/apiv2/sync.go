package apiv2

import (
	"vm-backend/internal/core/domain/sync"

	"github.com/gofiber/fiber/v2"
)

func RouteSync(router fiber.Router, endpoint sync.HTTPTransport) {
	api := router.Group("sync/:machine_id")

	api.Post("fetch", endpoint.FetchMachine)
	api.Post("push", endpoint.PushMachine)
	api.Post("slots/fetch", endpoint.FetchSlots)
	api.Post("slots/push", endpoint.PushSlots)
	api.Post("transactions/pull", endpoint.PullTransactions)
	api.Post("channels/fetch", endpoint.FetchChannels)
	api.Post("channels/push", endpoint.PushChannels)
}
