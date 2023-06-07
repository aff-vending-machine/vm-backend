package sync_http

import (
	"vm-backend/internal/core/infra/network/fiber/http"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) PushChannels(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeSyncRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.PushChannels(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}
