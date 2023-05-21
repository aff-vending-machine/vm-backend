package account_role

import "vm-backend/internal/core/domain/account"

type transportImpl struct {
	usecase account.RoleUsecase
}

func NewTransport(uc account.RoleUsecase) account.RoleTransport {
	return &transportImpl{uc}
}
