package rabbitmq

import (
	"sync"
	"vm-backend/internal/core/infra/network/rabbitmq"
)

type routerImpl struct {
	*rabbitmq.Service
	*rabbitmq.Server
	mu sync.Mutex
}

func New(service *rabbitmq.Service) *routerImpl {
	return &routerImpl{
		service,
		rabbitmq.NewServer(service.Connection),
		sync.Mutex{},
	}
}
