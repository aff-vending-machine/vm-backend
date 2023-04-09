package app

import (
	"github.com/aff-vending-machine/vm-backend/config"
	"github.com/aff-vending-machine/vm-backend/internal/boot/preload"
	"github.com/aff-vending-machine/vm-backend/internal/boot/registry"
	"github.com/aff-vending-machine/vm-backend/internal/boot/router/fiber"
	"github.com/aff-vending-machine/vm-backend/internal/boot/router/topic"
	"github.com/rs/zerolog/log"
)

func Run(cfg config.BootConfig) {
	log.Debug().Msg("init application")

	var (
		module    = NewModule(cfg)
		service   = NewService(module)
		usecase   = registry.NewUsecase(service)
		transport = NewTransport(usecase)
	)

	if cfg.App.Preload {
		superadmin := preload.CreateSuperAdminRole(usecase.Role)
		admin := preload.CreateAdminRole(usecase.Role)
		preload.CreateManagerRole(usecase.Role)
		preload.CreateStaffRole(usecase.Role)

		preload.CreateSuperAdmin(usecase.User, superadmin)
		preload.CreateAdmin(usecase.User, admin)
	}

	fiber.New(module.Fiber).Serve(transport.HTTP)
	topic.New(module.RabbitMQ).Serve(cfg.RabbitMQ.Queue, transport.Topic)

	log.Debug().Msg("start application")
}
