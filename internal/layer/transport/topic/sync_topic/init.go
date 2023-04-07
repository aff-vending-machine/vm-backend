package sync_topic

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase"
)

type syncImpl struct {
	usecase usecase.Sync
}

func New(uc usecase.Sync) *syncImpl {
	return &syncImpl{uc}
}
