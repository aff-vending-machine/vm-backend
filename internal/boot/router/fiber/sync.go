package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeSync(router fiber.Router, endpoint http.Sync) {
	router.Post("sync/:machine_id/fetch", endpoint.FetchMachine)
	router.Post("sync/:machine_id/push", endpoint.PushMachine)
	router.Post("sync/:machine_id/slots/fetch", endpoint.FetchSlots)
	router.Post("sync/:machine_id/slots/push", endpoint.PushSlots)
	router.Post("sync/:machine_id/transactions/pull", endpoint.PullTransactions)
}
