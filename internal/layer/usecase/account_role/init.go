package account_role

import (
	"vm-backend/internal/core/domain/account"
)

type usecaseImpl struct {
	roleRepo account.RoleRepository
}

func NewUsecase(
	arr account.RoleRepository,
) account.RoleUsecase {
	return &usecaseImpl{
		arr,
	}
}
