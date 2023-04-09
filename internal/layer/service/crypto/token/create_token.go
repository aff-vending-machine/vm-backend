package token

import (
	"context"
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/jwt"
	"github.com/aff-vending-machine/vm-backend/pkg/utils"
)

// CreateAccessToken creates an access token for the specified user.
func (m *managerImpl) CreateAccessToken(ctx context.Context, data jwt.Token) (string, error) {
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

	return createToken(claims, m.SecretAccess)
}

// CreateRefereshToken creates an refresh token for the specified user.
func (m *managerImpl) CreateRefreshToken(ctx context.Context, data jwt.Token) (string, error) {
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

	return createToken(claims, m.SecretRefresh)
}
