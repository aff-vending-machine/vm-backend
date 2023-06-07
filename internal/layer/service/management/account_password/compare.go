package account_password

import (
	"context"

	"vm-backend/pkg/helpers/errs"

	"golang.org/x/crypto/bcrypt"
)

func (m *managementImpl) Compare(ctx context.Context, hash string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if errs.IsErr(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
