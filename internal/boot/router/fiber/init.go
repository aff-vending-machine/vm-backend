package fiber

import (
	"sync"
	"vm-backend/internal/core/infrastructure/network/fiber"
)

type routerImpl struct {
	*fiber.Server
	mu sync.Mutex
}

func New(server *fiber.Server) *routerImpl {
	return &routerImpl{
		server,
		sync.Mutex{},
	}
}