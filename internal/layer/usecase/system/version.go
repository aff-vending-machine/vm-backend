package system

import (
	"context"

	"vm-backend/internal/layer/usecase/system/response"
)

func (uc *usecaseImpl) Version(ctx context.Context) (*response.Version, error) {
	return &response.Version{Version: "2.0.0"}, nil
}
