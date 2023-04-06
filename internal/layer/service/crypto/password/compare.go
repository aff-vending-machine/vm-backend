package password

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

func (m *managerImpl) Compare(ctx context.Context, hash string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if errs.IsErr(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
