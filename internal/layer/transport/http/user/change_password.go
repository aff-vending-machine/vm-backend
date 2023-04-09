package user

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/request"
	"github.com/gofiber/fiber/v2"
)

func (r *restImpl) ChangePassword(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeChangePasswordRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.ChangePassword(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}

func makeChangePasswordRequest(c *fiber.Ctx) (*request.ChangePassword, error) {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	var req request.ChangePassword
	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}
	req.ID = uint(id)

	return &req, nil
}
