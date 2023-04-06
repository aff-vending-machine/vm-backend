package auth_http

import (
	"strings"

	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/request"
	"github.com/gofiber/fiber/v2"
)

func (r *restImpl) PermissionRequired(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeGetPermissionLevelRequest(c)
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
	case 0:
		if c.Method() == fiber.MethodGet {
			return c.Next()
		}

	case 1:
		if c.Method() == fiber.MethodGet {
			return c.Next()
		} else if len(paths) >= 4 && paths[4] == "me" {
			return c.Next()
		}

	case 2:
		return c.Next()

	case 3:
		return c.Next()

	}

	err = fiber.ErrForbidden
	return http.Forbidden(c, err)
}

func makeGetPermissionLevelRequest(c *fiber.Ctx) (*request.GetPermissionLevel, error) {
	paths := strings.Split(c.Path(), "/")
	skip := whitelist(c, paths)
	if skip {
		return nil, nil
	}

	selection := paths[3]
	scope := selection[:len(selection)-1]

	userID, err := getUserID(c)
	if err != nil {
		return nil, err
	}

	return &request.GetPermissionLevel{
		UserID: userID,
		Scope:  scope,
	}, nil
}
