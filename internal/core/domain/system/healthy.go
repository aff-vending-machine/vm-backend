package system

import (
	"context"

	"vm-backend/internal/layer/usecase/system/response"
)

type Usecase interface {
	Healthy(context.Context) (*response.Healthy, error)
	Version(context.Context) (*response.Version, error)
}
