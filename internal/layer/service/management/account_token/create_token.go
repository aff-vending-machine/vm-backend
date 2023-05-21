package account_token

import (
	"context"
	"time"

	"vm-backend/internal/core/domain/account"
	"vm-backend/pkg/utils"
)

// CreateAccessToken creates an access token for the specified user.
func (m *managementImpl) CreateAccessToken(ctx context.Context, data account.Token) (string, error) {
	// Create the JWT claims.
	claims := MapClaims{
		"jti":  utils.GenerateRandom(8),           // claim to uniquely identify the JWT and prevent replay attacks.
		"sub":  data.ID,                           // identify the user
		"iat":  time.Now().Unix(),                 // claim to set the time at which the JWT was issued.
		"exp":  time.Now().Add(data.Alive).Unix(), // claim to identify the user or entity that the JWT is issued to.
		"aud":  m.audience,                        // claim to identify the intended recipient of the JWT.
		"iss":  m.issuer,                          // claim to identify the issuer of the JWT.
		"type": data.Type,
		"name": data.Name,
		"role": data.Role,
	}

	return createToken(claims, m.secretAccess)
}

// CreateRefereshToken creates an refresh token for the specified user.
func (m *managementImpl) CreateRefreshToken(ctx context.Context, data account.Token) (string, error) {
	// Create the JWT claims.
	claims := MapClaims{
		"jti":  utils.GenerateRandom(13),          // claim to uniquely identify the JWT and prevent replay attacks.
		"sub":  data.ID,                           // identify the user
		"iat":  time.Now().Unix(),                 // claim to set the time at which the JWT was issued.
		"exp":  time.Now().Add(data.Alive).Unix(), // claim to identify the user or entity that the JWT is issued to.
		"aud":  m.audience,                        // claim to identify the intended recipient of the JWT.
		"iss":  m.issuer,                          // claim to identify the issuer of the JWT.
		"azp":  m.authorizedParty,                 // claim to identify the party authorized to generate access tokens using the refresh token.
		"type": data.Type,
		"name": data.Name,
		"role": data.Role,
	}

	return createToken(claims, m.secretRefresh)
}
