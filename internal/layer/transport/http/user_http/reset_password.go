package user_http

import (
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/request"
	"github.com/gofiber/fiber/v2"
)

func (r *restImpl) ResetPassword(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeResetPasswordRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.ResetPassword(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}

func makeResetPasswordRequest(c *fiber.Ctx) (*request.ResetPassword, error) {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	if c.Locals("level") == nil {
		return &request.ResetPassword{ID: uint(id), Level: 0}, nil
	}
	local := c.Locals("level")

	level := 0
	switch lvl := local.(type) {
	case int:
		level = lvl
	case uint:
		level = int(lvl)
	case float64:
		level = int(lvl)
	default:
		return nil, fmt.Errorf("level is not number type: %T", local)
	}

	return &request.ResetPassword{ID: uint(id), Level: level}, nil
}
