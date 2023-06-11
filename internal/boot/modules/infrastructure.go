package modules

import (
	"vm-backend/configs"
	"vm-backend/internal/core/infra/network/fiber"
	"vm-backend/internal/core/infra/network/rabbitmq"
	"vm-backend/internal/core/infra/network/websocket"
	"vm-backend/internal/core/infra/strorage/postgresql"
)

type Infrastructure struct {
	App        configs.AppConfig
	Fiber      *fiber.Server
	PostgreSQL *postgresql.Client
	RabbitMQ   *rabbitmq.Service
	WebSocket  *websocket.Server
}
