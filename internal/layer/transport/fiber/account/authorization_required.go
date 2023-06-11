package account

import (
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/core/infra/network/fiber/http"
	"vm-backend/internal/layer/usecase/account/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) AuthorizationRequired(c *fiber.Ctx) error {
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
	c.Locals(account.AccessIDKey, claims.UserID)
	c.Locals(account.AccessKey, claims.Username)
	c.Locals(account.RoleIDKey, claims.RoleID)
	c.Locals(account.RoleKey, claims.Role)
	c.Locals(account.BranchIDKey, claims.BranchID)
	c.Locals(account.BranchKey, claims.Branch)

	return c.Next()
}

func makeValidateTokenRequest(c *fiber.Ctx) (*request.ValidateToken, error) {
	token, err := getBearerAuthorization(c)
	if err != nil {
		return nil, err
	}

	return &request.ValidateToken{Token: token}, nil
}
