package auth

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/request"
	"github.com/gofiber/fiber/v2"
)

func (r *httpImpl) RefreshToken(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeValidateRefreshTokenRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	res, err := r.usecase.ValidateRefreshToken(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}

func makeValidateRefreshTokenRequest(c *fiber.Ctx) (*request.ValidateToken, error) {
	token, err := getBearerAuthorization(c)
	if err != nil {
		return nil, err
	}

	req := &request.ValidateToken{Token: token}
	return req, nil
}
