package response

import "github.com/aff-vending-machine/vm-backend/internal/core/domain/jwt"

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	Type   string `json:"type"`
}

func ToToken(c *jwt.Token) *Claims {
	return &Claims{
		UserID: c.ID,
		Role:   c.Role,
		Type:   c.Type,
	}
}
