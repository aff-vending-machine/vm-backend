package token

import (
	"context"
	"encoding/json"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/jwt"
)

func (m *managerImpl) ValidateAccessToken(ctx context.Context, t string) (*jwt.Token, error) {
	claims, err := m.validateToken(t, m.SecretAccess, "ACCESS_TOKEN")
	if err != nil {
		return nil, err
	}

	var result jwt.Token
	b, _ := json.Marshal(claims)
	json.Unmarshal(b, &result)

	return &result, nil
}

func (m *managerImpl) ValidateRefreshToken(ctx context.Context, t string) (*jwt.Token, error) {
	claims, err := m.validateToken(t, m.SecretRefresh, "REFRESH_TOKEN")
	if err != nil {
		return nil, err
	}

	var result jwt.Token
	b, _ := json.Marshal(claims)
	json.Unmarshal(b, &result)

	return &result, nil
}
