package auth

import (
	"strings"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/permission"
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/request"
	"github.com/aff-vending-machine/vm-backend/pkg/trace"
	"github.com/gofiber/fiber/v2"
)

func (r *httpImpl) PermissionRequired(c *fiber.Ctx) error {
	ctx, span := trace.Start(c.Context())
	defer span.End()

	req, err := makeGetPermissionLevel(c)
	if err != nil {
		trace.RecordError(span, err)
		return http.Forbidden(c, err)
	}
	if req == nil && err == nil {
		return c.Next()
	}

	// usecase execution
	result, err := r.usecase.GetPermissionLevel(ctx, req)
	if err != nil {
		trace.RecordError(span, err)
		return http.UsecaseError(c, err)
	}

	paths := strings.Split(c.Path(), "/")
	c.Locals("x-permission-level", result.Level)

	switch result.Level {
	case permission.Viewer:
		if c.Method() == fiber.MethodGet {
			return c.Next()
		}

	case permission.Owner:
		if c.Method() == fiber.MethodGet {
			return c.Next()
		} else if len(paths) >= 4 && paths[4] == "me" {
			return c.Next()
		}

	case permission.Editor:
		return c.Next()

	case permission.Admin:
		return c.Next()

	}

	err = fiber.ErrForbidden
	trace.RecordError(span, err)
	return http.Forbidden(c, err)
}

func makeGetPermissionLevel(c *fiber.Ctx) (*request.GetPermissionLevel, error) {
	paths := strings.Split(c.Path(), "/")
	skip := whitelist(c, paths)
	if skip {
		return nil, nil
	}

	selection := paths[3]
	scope, _ := strings.CutSuffix(selection, "s")

	userID, err := getUserID(c)
	if err != nil {
		return nil, err
	}

	return &request.GetPermissionLevel{
		UserID: userID,
		Scope:  scope,
	}, nil
}
