package jwt

import (
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Token struct {
	ID    uint          `json:"sub"`
	Name  string        `json:"name"`
	Role  string        `json:"role"`
	Type  string        `json:"type"`
	Alive time.Duration `json:"-,omitempty"`
}

func NewAccessToken(user entity.User) Token {
	return Token{
		ID:    user.ID,
		Name:  user.Username,
		Role:  user.Role.Name,
		Type:  "ACCESS_TOKEN",
		Alive: 1 * time.Hour,
	}
}

func NewRefreshToken(user entity.User) Token {
	return Token{
		ID:    user.ID,
		Name:  user.Username,
		Role:  user.Role.Name,
		Type:  "REFRESH_TOKEN",
		Alive: 24 * time.Hour,
	}
}
