package system_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/system/response"
)

func (uc *usecaseImpl) Version(ctx context.Context) (*response.Version, error) {
	return &response.Version{Version: "1.0.0"}, nil
}
