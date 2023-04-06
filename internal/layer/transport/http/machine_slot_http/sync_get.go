package machine_slot_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/gofiber/fiber/v2"
)

func (r *restImpl) SyncGet(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeSyncRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.SyncGet(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}
