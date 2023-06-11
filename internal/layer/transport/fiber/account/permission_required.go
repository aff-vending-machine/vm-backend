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
	result, err := r.usecase.GetAccountPermission(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	paths := strings.Split(c.Path(), "/")
	c.Locals(account.PermissionLevelKey, result.Level)
	c.Locals(account.AccessScopeKey, req.Scope)

	switch result.Level {
	case account.Viewer:
		if c.Method() == fiber.MethodGet {
			c.Locals(account.AccessConditionKey, "viewer-get-method")
			return c.Next()
		}
		c.Locals(account.AccessConditionKey, "viewer-no-permission")

	case account.Owner:
		if c.Method() == fiber.MethodGet {
			c.Locals(account.AccessConditionKey, "owner-get-method")
			return c.Next()
		} else if len(paths) >= 5 && paths[4] == "me" {
			c.Locals(account.AccessConditionKey, "owner-me-path")
			return c.Next()
		} else if result.BranchID == 0 {
			c.Locals(account.AccessConditionKey, "owner-branch-all")
			return c.Next()
		} else if branchID, ok := c.Locals(account.BranchIDKey).(uint); ok {
			if branchID == result.BranchID {
				c.Locals(account.AccessConditionKey, "owner-branch-match")
				return c.Next()
			}
		}
		c.Locals(account.AccessConditionKey, "owner-no-permission")

	case account.Editor:
		c.Locals(account.AccessConditionKey, "editor")
		return c.Next()

	case account.Admin:
		c.Locals(account.AccessConditionKey, "admin")
		return c.Next()

	}

	err = fiber.ErrForbidden
	return http.Forbidden(c, err)
}

func makeGetPermissionLevel(c *fiber.Ctx) (*request.GetAccountPermission, error) {
	paths := strings.Split(c.Path(), "/")
	skip := whitelist(c, paths)
	if skip {
		return nil, nil
	}

	scope := paths[3]
	userID, err := account.GetAccessID(c)
	if err != nil {
		return nil, err
	}

	return &request.GetAccountPermission{
		UserID: userID,
		Scope:  scope,
	}, nil
}
