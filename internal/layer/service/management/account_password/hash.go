package account_password

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (m *managementImpl) HashPassword(ctx context.Context, password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), m.salt)

	if err != nil {
		return "", err
	}

	result := string(hashed)
	return result, nil
}
