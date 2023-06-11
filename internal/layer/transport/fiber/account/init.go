package account

import (
	"fmt"
	"strings"
	"vm-backend/internal/core/domain/account"

	"github.com/gofiber/fiber/v2"
)

type transportImpl struct {
	usecase account.Usecase
}

func NewTransport(uc account.Usecase) account.Transport {
	return &transportImpl{uc}
}

func whitelist(c *fiber.Ctx, paths []string) bool {
	if len(paths) <= 4 {
		return false
	}

	if c.Method() == fiber.MethodGet {
		// GET /systems/dashboard
		if paths[3] == "systems" && paths[4] == "dashboard" {
			return true
		}
	}

	if c.Method() == fiber.MethodPost {
		// POST /users/me/change-password
		if len(paths) == 6 && paths[3] == "users" && paths[4] == "me" && paths[5] == "change-password" {
			return true
		}
	}

	return false
}

func getBearerAuthorization(ctx *fiber.Ctx) (string, error) {
	authHeader := ctx.Get(fiber.HeaderAuthorization)
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is missing")
	}

	authFields := strings.Fields(authHeader)
	if len(authFields) != 2 || strings.ToLower(authFields[0]) != "bearer" {
		return "", fmt.Errorf("authorization header is malformed")
	}

	return authFields[1], nil
}
