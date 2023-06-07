package account_user

import (
	"vm-backend/internal/core/infra/network/fiber/http"
	"vm-backend/internal/layer/usecase/account_user/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) ChangePassword(c *fiber.Ctx) error {
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
	var req request.ChangePassword
	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}
	req.ID = uint(http.Int(c, "x-access-id"))

	return &req, nil
}
