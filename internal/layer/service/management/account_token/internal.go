package account_token

import (
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	jtiCache  = map[string]struct{}{}
	cacheLock sync.Mutex
)

type MapClaims = map[string]interface{}

// createToken creates an token for the specified claims.
func createToken(claims jwt.MapClaims, secret string) (string, error) {

	// Create the JWT token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and return the JWT.
	return token.SignedString([]byte(secret))
}

func (m *managementImpl) validateToken(token string, secret string, tokenType string) (jwt.MapClaims, error) {
	claims, err := parseToken(token, secret)
	if err != nil {
		return nil, err
	}

	err = m.validateClaims(claims, false)
	if err != nil {
		return nil, err
	}

	// Get the JWT type claim.
	tt, ok := claims["type"].(string)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}

	// Validate the JWT type claim.
	if tt != tokenType {
		return nil, jwt.ErrInvalidKeyType
	}

	if tokenType != "refresh_token" {
		return claims, nil
	}

	// Get the JWT authorized party claim.
	azp, ok := claims["azp"].(string)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}

	// Validate the JWT authorized party claim.
	if err := m.validateAzp(azp); err != nil {
		return nil, err
	}

	return claims, nil
}

func parseToken(token, secret string) (jwt.MapClaims, error) {
	// Parse the signed JWT with the secret key.
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	// Get the claims from the parsed token.
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}

// validateClaims checks the JWT Token is valid.
func (m *managementImpl) validateClaims(claims jwt.MapClaims, checkJTI bool) error {
	// Get the JWT ID.
	jti, ok := claims["jti"].(string)
	if !ok {
		return jwt.ErrInvalidKey
	}

	// Validate the JWT ID.
	if err := validateJti(jti); checkJTI && err != nil {
		return err
	}

	// Get the JWT audience claims.
	aud, ok := claims["aud"].(string)
	if !ok {
		return jwt.ErrInvalidKey
	}

	// Validate the JWT audience claims.
	if err := m.validateAudience(aud); err != nil {
		return err
	}

	// Get the JWT issuer claims.
	iss, ok := claims["iss"].(string)
	if !ok {
		return jwt.ErrInvalidKey
	}

	// Validate the JWT issuer claims.
	if err := m.validateIssuer(iss); err != nil {
		return err
	}

	// Get the JWT issued at claims.
	iat, ok := claims["iat"].(float64)
	if !ok {
		return jwt.ErrInvalidKey
	}

	// Validate the JWT issued at claims.
	if err := validateIssuedAt(iat); err != nil {
		return err
	}

	// Get the JWT expiration claims.
	exp, ok := claims["exp"].(float64)
	if !ok {
		return jwt.ErrInvalidKey
	}

	// Validate the JWT issued at claims.
	if err := validateExpiration(exp); err != nil {
		return err
	}

	return nil
}

// validateJti checks whether the JWT ID is valid and has not been used before.
func validateJti(jti string) error {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	if _, ok := jtiCache[jti]; ok {
		// The JWT ID has been used before.
		return jwt.ErrTokenInvalidId
	}

	// The JWT ID is valid and has not been used before.
	jtiCache[jti] = struct{}{}
	return nil
}

// validateAudience checks whether the JWT audience claims is valid.
func (m *managementImpl) validateAudience(aud string) error {
	if aud != m.audience {
		// The JWT audience claim is invalid.
		return jwt.ErrTokenInvalidAudience
	}

	// The JWT audience claims are valid.
	return nil
}

// validateAudience checks whether the JWT issuer claims is valid.
func (m *managementImpl) validateIssuer(iss string) error {
	if iss != m.issuer {
		// The JWT issuer claim is invalid.
		return jwt.ErrTokenInvalidIssuer
	}

	// The JWT issuer claims are valid.
	return nil
}

// validateIssuedAt checks whether the JWT issued at claims are valid.
func validateIssuedAt(iat float64) error {
	issuedAt := time.Unix(int64(iat), 0)
	now := time.Now()
	if issuedAt.After(now) {
		// The JWT was issued in the future.
		return jwt.ErrTokenUsedBeforeIssued
	}

	// The JWT issued at claims are valid.
	return nil
}

// validateExpiration checks whether the JWT expiration claims are valid.
func validateExpiration(exp float64) error {
	expiration := time.Unix(int64(exp), 0)
	now := time.Now()
	if expiration.Before(now) {
		// The JWT has expired.
		return jwt.ErrTokenExpired
	}

	// The JWT issued at and expiration claims are valid.
	return nil
}

// validateAzp checks whether the JWT authorized party claim is valid.
func (m *managementImpl) validateAzp(azp string) error {
	if azp != m.authorizedParty {
		// The JWT authorized party claim is invalid.
		return jwt.ErrTokenMalformed
	}

	// The JWT authorized party claim is valid.
	return nil
}
