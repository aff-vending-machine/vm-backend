package system

import "vm-backend/internal/core/domain/system"

type usecaseImpl struct {
}

func NewUsecase() system.Usecase {
	return &usecaseImpl{}
}
