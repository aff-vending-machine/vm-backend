package app

import (
	"github.com/aff-vending-machine/vm-backend/config"
	"github.com/aff-vending-machine/vm-backend/internal/boot/registry"
	"github.com/aff-vending-machine/vm-backend/internal/boot/router/fiber"
	"github.com/rs/zerolog/log"
)

func Run(cfg config.BootConfig) {
	var (
		module    = NewModule(cfg)
		service   = NewService(module)
		usecase   = registry.NewUsecase(service)
		transport = NewTransport(usecase)
	)

	fiber.New(module.Fiber).Serve(transport.HTTP)

	log.Debug().Msg("start application")
}
