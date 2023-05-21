package catalog_product

import (
	"vm-backend/internal/core/domain/catalog"
)

type restImpl struct {
	usecase catalog.ProductUsecase
}

func NewTransport(uc catalog.ProductUsecase) catalog.ProductTransport {
	return &restImpl{uc}
}
