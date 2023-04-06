package auth_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/request"
	"github.com/gofiber/fiber/v2"
)

func (r *restImpl) AuthorizationRequired(c *fiber.Ctx) error {
	ctx := c.UserContext()

	// get token from request header
	req, err := makeValidateTokenRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	claims, err := r.usecase.ValidateAccessToken(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	// save session data in context
	c.Locals("x-access", claims.UserID)

	return c.Next()
}

func makeValidateTokenRequest(c *fiber.Ctx) (*request.ValidateToken, error) {
	token, err := getBearerAuthorization(c)
	if err != nil {
		return nil, err
	}

	return &request.ValidateToken{Token: token}, nil
}
