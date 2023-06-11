package store_branch

import (
	"vm-backend/internal/core/domain/store"
)

type restImpl struct {
	usecase store.BranchUsecase
}

func NewTransport(uc store.BranchUsecase) store.BranchTransport {
	return &restImpl{uc}
}
