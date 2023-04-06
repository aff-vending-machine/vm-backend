package crypto

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/jwt"
)

type Password interface {
	HashPassword(ctx context.Context, password string) (string, error)
	Compare(ctx context.Context, hashed string, password string) (bool, error)
}

type Token interface {
	CreateAccessToken(context.Context, jwt.Token) (string, error)
	CreateRefreshToken(context.Context, jwt.Token) (string, error)
	ValidateAccessToken(context.Context, string) (*jwt.Token, error)
	ValidateRefreshToken(context.Context, string) (*jwt.Token, error)
}
