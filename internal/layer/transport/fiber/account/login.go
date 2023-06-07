package account

import (
	"vm-backend/internal/core/infra/network/fiber/http"
	"vm-backend/internal/layer/usecase/account/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) Login(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeLoginRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	res, err := r.usecase.Login(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}

func makeLoginRequest(c *fiber.Ctx) (*request.Login, error) {
	var req request.Login

	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
