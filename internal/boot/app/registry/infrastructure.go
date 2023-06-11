package registry

import (
	"vm-backend/configs"
	"vm-backend/internal/boot/modules"
	"vm-backend/internal/core/infra/network/fiber"
	"vm-backend/internal/core/infra/network/rabbitmq"
	"vm-backend/internal/core/infra/network/websocket"
	"vm-backend/internal/core/infra/strorage/postgresql"
)

func NewInfrastructure(cfg configs.Config) modules.Infrastructure {
	return modules.Infrastructure{
		App:        cfg.App,
		Fiber:      fiber.New(cfg.Fiber),
		PostgreSQL: postgresql.New(cfg.PostgreSQL),
		RabbitMQ:   rabbitmq.New(cfg.RabbitMQ),
		WebSocket:  websocket.New(cfg.WebSocket),
	}
}
