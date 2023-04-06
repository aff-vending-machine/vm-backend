package product_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/product"
)

type restImpl struct {
	usecase product.Usecase
}

func New(uc product.Usecase) *restImpl {
	return &restImpl{uc}
}
