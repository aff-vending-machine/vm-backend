package role_usecase

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

type usecaseImpl struct {
	roleRepo repository.Role
}

func New(r repository.Role) *usecaseImpl {
	return &usecaseImpl{r}
}
