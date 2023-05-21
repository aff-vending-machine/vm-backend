package account_user

import (
	"vm-backend/internal/core/domain/account"
)

type transportImpl struct {
	usecase account.UserUsecase
}

func NewTransport(uc account.UserUsecase) account.UserTransport {
	return &transportImpl{uc}
}
