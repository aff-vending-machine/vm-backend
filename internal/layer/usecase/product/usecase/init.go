package product_usecase

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

type usecaseImpl struct {
	productRepo repository.Product
}

func New(p repository.Product) *usecaseImpl {
	return &usecaseImpl{p}
}
