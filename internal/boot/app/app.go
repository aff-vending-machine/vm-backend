package app

import (
	"vm-backend/configs"
	"vm-backend/internal/boot/app/registry"
	"vm-backend/internal/boot/migration"
	"vm-backend/internal/boot/router/fiber"
	"vm-backend/internal/boot/router/rabbitmq"

	"github.com/rs/zerolog/log"
)

func Run(cfg configs.Config) {
	log.Debug().Msg("init application")

	var (
		infra     = registry.NewInfrastructure(cfg)
		service   = registry.NewService(infra)
		usecase   = registry.NewUsecase(service)
		transport = registry.NewTransport(usecase)
	)

	if cfg.App.Migration {
		log.Info().Msg("migration")
		migration.CreateBranchFromMachine(service.Repository)
		migration.UpdateAlternativeScope(service.Repository)
		migration.UpdateBranchInTransactions(service.Repository)
	}

	fiber.New(infra.Fiber).Serve(transport.Fiber)
	rabbitmq.New(infra.RabbitMQ).Serve(transport.RabbitMQ)

	log.Debug().Msg("start application")
}
