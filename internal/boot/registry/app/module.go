package app

import (
	"github.com/aff-vending-machine/vm-backend/config"
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber"
	"github.com/aff-vending-machine/vm-backend/internal/core/module/postgresql"
	"github.com/aff-vending-machine/vm-backend/internal/core/module/rabbitmq"
)

// Infrastructure
type Module struct {
	Config     config.BootConfig
	Fiber      *fiber.Wrapper
	PostgreSQL *postgresql.Wrapper
	RabbitMQ   *rabbitmq.Wrapper
}

func NewModule(cfg config.BootConfig) Module {
	return Module{
		Config:     cfg,
		Fiber:      fiber.New(cfg.Fiber),
		PostgreSQL: postgresql.New(cfg.PostgreSQL),
		RabbitMQ:   rabbitmq.New(cfg.RabbitMQ),
	}
}
