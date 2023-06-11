package account

import (
	"strings"
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/core/infra/network/fiber/http"
	"vm-backend/internal/layer/usecase/account/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) PermissionRequired(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeGetPermissionLevel(c)
	if err != nil {
		return http.Forbidden(c, err)
	}
	if req == nil && err == nil {
		return c.Next()
	}

	// usecase execution
	result, err := r.usecase.GetPermissionLevel(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	paths := strings.Split(c.Path(), "/")
	c.Locals("x-permission-level", result.Level)

	switch result.Level {
	case account.Viewer:
		if c.Method() == fiber.MethodGet {
			return c.Next()
		}

	case account.Owner:
		if c.Method() == fiber.MethodGet {
			return c.Next()
		} else if len(paths) >= 4 && paths[4] == "me" {
			return c.Next()
		}

	case account.Editor:
		return c.Next()

	case account.Admin:
		return c.Next()

	}

	err = fiber.ErrForbidden
	return http.Forbidden(c, err)
}

func makeGetPermissionLevel(c *fiber.Ctx) (*request.GetPermissionLevel, error) {
	paths := strings.Split(c.Path(), "/")
	skip := whitelist(c, paths)
	if skip {
		return nil, nil
	}

	scope := paths[3]
	userID, err := getUserID(c)
	if err != nil {
		return nil, err
	}

	return &request.GetPermissionLevel{
		UserID: userID,
		Scope:  scope,
	}, nil
}
