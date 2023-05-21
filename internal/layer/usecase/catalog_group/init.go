package catalog_group

import (
	"vm-backend/internal/core/domain/catalog"
)

type usecaseImpl struct {
	catalogGroupRepo catalog.GroupRepository
}

func NewUsecase(
	cpr catalog.GroupRepository,
) catalog.GroupUsecase {
	return &usecaseImpl{
		cpr,
	}
}
