package catalog_group

import (
	"vm-backend/internal/core/domain/catalog"
)

type restImpl struct {
	usecase catalog.GroupUsecase
}

func NewTransport(uc catalog.GroupUsecase) catalog.GroupTransport {
	return &restImpl{uc}
}
