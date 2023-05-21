package account_role

import (
	"vm-backend/internal/core/infrastructure/network/fiber/http"
	"vm-backend/internal/layer/usecase/account_role/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeDeleteRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.Delete(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}

func makeDeleteRequest(c *fiber.Ctx) (*request.Delete, error) {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	return &request.Delete{ID: uint(id)}, nil
}
