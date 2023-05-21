package account

import (
	"context"
	"vm-backend/internal/layer/usecase/account/request"
	"vm-backend/internal/layer/usecase/account/response"

	"github.com/gofiber/fiber/v2"
)

type Usecase interface {
	Login(ctx context.Context, req *request.Login) (*response.AuthResult, error)
	ValidateAccessToken(ctx context.Context, req *request.ValidateToken) (*response.Claims, error)
	ValidateRefreshToken(ctx context.Context, req *request.ValidateToken) (*response.AuthResult, error)
	GetPermissionLevel(ctx context.Context, req *request.GetPermissionLevel) (*response.PermissionLevel, error)
}

type PasswordManagement interface {
	HashPassword(ctx context.Context, password string) (string, error)
	Compare(ctx context.Context, hashed string, password string) (bool, error)
}

type TokenManagement interface {
	CreateAccessToken(context.Context, Token) (string, error)
	CreateRefreshToken(context.Context, Token) (string, error)
	ValidateAccessToken(context.Context, string) (*Token, error)
	ValidateRefreshToken(context.Context, string) (*Token, error)
}

type Transport interface {
	Login(ctx *fiber.Ctx) error                 // POST {auth/login}
	RefreshToken(ctx *fiber.Ctx) error          // POST {auth/refresh-token}
	AuthorizationRequired(ctx *fiber.Ctx) error // middleware
	PermissionRequired(ctx *fiber.Ctx) error    // middleware
}
