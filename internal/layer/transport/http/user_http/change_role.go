package user_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/request"
	"github.com/gofiber/fiber/v2"
)

func (r *restImpl) ChangeRole(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeChangeRoleRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.ChangeRole(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}

func makeChangeRoleRequest(c *fiber.Ctx) (*request.ChangeRole, error) {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	var req request.ChangeRole
	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}
	req.ID = uint(id)

	return &req, nil
}
