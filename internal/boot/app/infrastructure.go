package app

import (
	"vm-backend/configs"
	"vm-backend/internal/core/infrastructure/network/fiber"
	"vm-backend/internal/core/infrastructure/network/rabbitmq"
	"vm-backend/internal/core/infrastructure/network/websocket"
	"vm-backend/internal/core/infrastructure/strorage/postgresql"
)

type Infrastructure struct {
	App        configs.AppConfig
	Fiber      *fiber.Server
	PostgreSQL *postgresql.Client
	RabbitMQ   *rabbitmq.Service
	WebSocket  *websocket.Server
}

func NewInfrastructure(cfg configs.Config) Infrastructure {
	return Infrastructure{
		App:        cfg.App,
		Fiber:      fiber.New(cfg.Fiber),
		PostgreSQL: postgresql.New(cfg.PostgreSQL),
		RabbitMQ:   rabbitmq.New(cfg.RabbitMQ),
		WebSocket:  websocket.New(cfg.WebSocket),
	}
}
