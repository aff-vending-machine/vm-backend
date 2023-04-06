package auth

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/response"
)

type Usecase interface {
	Login(ctx context.Context, req *request.Login) (*response.AuthResult, error)
	ValidateAccessToken(ctx context.Context, req *request.ValidateToken) (*response.Claims, error)
	ValidateRefreshToken(ctx context.Context, req *request.ValidateToken) (*response.AuthResult, error)
	GetPermissionLevel(ctx context.Context, req *request.GetPermissionLevel) (*response.PermissionLevel, error)
}
