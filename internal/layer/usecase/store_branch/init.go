package store_branch

import (
	"vm-backend/internal/core/domain/store"
)

type usecaseImpl struct {
	storeBranchRepo store.BranchRepository
}

func NewUsecase(
	sbr store.BranchRepository,
) store.BranchUsecase {
	return &usecaseImpl{
		sbr,
	}
}
