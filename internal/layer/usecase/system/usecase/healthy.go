package system_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/system/response"
)

func (uc *usecaseImpl) Healthy(ctx context.Context) (*response.Healthy, error) {
	return &response.Healthy{
		Ready:   true,
		Message: "OK",
	}, nil
}
