package fiber

import (
	"vm-backend/internal/core/infra/network/fiber"
)

type routerImpl struct {
	*fiber.Server
}

func New(server *fiber.Server) *routerImpl {
	return &routerImpl{
		server,
	}
}
