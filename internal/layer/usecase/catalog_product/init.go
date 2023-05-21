package catalog_product

import (
	"vm-backend/internal/core/domain/catalog"
)

type usecaseImpl struct {
	productRepo catalog.ProductRepository
}

func NewUsecase(
	cpr catalog.ProductRepository,
) catalog.ProductUsecase {
	return &usecaseImpl{
		cpr,
	}
}
