package account_user

import (
	"vm-backend/internal/core/domain/account"
)

type usecaseImpl struct {
	userRepo    account.UserRepository
	roleRepo    account.RoleRepository
	passwordMgr account.PasswordManagement
}

func NewUsecase(
	aur account.UserRepository,
	arr account.RoleRepository,
	apm account.PasswordManagement,
) account.UserUsecase {
	return &usecaseImpl{
		aur,
		arr,
		apm,
	}
}
