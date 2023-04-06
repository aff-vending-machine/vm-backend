package auth_http

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func whitelist(c *fiber.Ctx, paths []string) bool {
	if len(paths) <= 4 {
		return false
	}

	if paths[3] != "users" && paths[4] != "me" {
		return false
	}

	if len(paths) == 6 && paths[5] == "change-password" && c.Method() == fiber.MethodPost {
		return true
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

func getUserID(ctx *fiber.Ctx) (uint, error) {
	if ctx.Locals("x-access") == nil {
		return 0, fmt.Errorf("user id is not exist")
	}

	str := fmt.Sprintf("%v", ctx.Locals("x-access"))
	id, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
