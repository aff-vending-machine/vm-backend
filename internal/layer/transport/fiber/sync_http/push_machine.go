package sync_http

import (
	"vm-backend/internal/core/infrastructure/network/fiber/http"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) PushMachine(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeSyncRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.PushMachine(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}