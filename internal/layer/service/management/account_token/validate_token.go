package account_token

import (
	"context"
	"vm-backend/internal/core/domain/account"
	"vm-backend/pkg/helpers/conv"
)

func (m *managementImpl) ValidateAccessToken(ctx context.Context, t string) (*account.Token, error) {
	claims, err := m.validateToken(t, m.secretAccess, "ACCESS_TOKEN")
	if err != nil {
		return nil, err
	}

	return conv.StructTo[account.Token](claims)
}

func (m *managementImpl) ValidateRefreshToken(ctx context.Context, t string) (*account.Token, error) {
	claims, err := m.validateToken(t, m.secretRefresh, "REFRESH_TOKEN")
	if err != nil {
		return nil, err
	}

	return conv.StructTo[account.Token](claims)
}
