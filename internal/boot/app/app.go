package app

import (
	"vm-backend/configs"
	"vm-backend/internal/boot/router/fiber"
	"vm-backend/internal/boot/router/rabbitmq"

	"github.com/rs/zerolog/log"
)

func Run(cfg configs.Config) {
	var (
		infra     = NewInfrastructure(cfg)
		service   = NewService(infra)
		usecase   = NewUsecase(service)
		transport = NewTransport(usecase)
	)

	fiber.New(infra.Fiber).Serve(transport.Fiber)
	rabbitmq.New(infra.RabbitMQ).Serve(transport.RabbitMQ)

	log.Debug().Msg("start application")
}
