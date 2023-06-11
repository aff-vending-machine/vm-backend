package modules

import (
	"vm-backend/internal/boot/router/fiber"
	"vm-backend/internal/boot/router/rabbitmq"
)

// Interface Adapter layers (driver)
type Transport struct {
	Fiber    fiber.Transport
	RabbitMQ rabbitmq.Transport
}
