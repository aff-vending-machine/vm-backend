package app

import (
	"vm-backend/configs"
	"vm-backend/internal/boot/migration"
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

	migration.MigrateUser(infra.PostgreSQL.DB)
	migration.MigrateRolePermission(infra.PostgreSQL.DB)
	migration.MigrateProduct(infra.PostgreSQL.DB)
	migration.MigrateTransaction(infra.PostgreSQL.DB)

	fiber.New(infra.Fiber).Serve(transport.Fiber)
	rabbitmq.New(infra.RabbitMQ).Serve(transport.RabbitMQ)

	log.Debug().Msg("start application")
}
