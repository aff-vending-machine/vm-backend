package fiber

import (
	"vm-backend/configs"
	"vm-backend/internal/core/infrastructure/network/fiber/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Server struct {
	*fiber.App
	configs.FiberConfig
}

func New(cfg configs.FiberConfig) *Server {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               cfg.Prefork,
		CaseSensitive:         cfg.CaseSensitive,
		StrictRouting:         cfg.StrictRouting,
		ServerHeader:          cfg.ServerHeader,
		AppName:               cfg.AppName,
	})

	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
		AllowHeaders:     "accept,content-type,authorization",
		AllowCredentials: true,
		MaxAge:           1728000,
	}))
	app.Use(middleware.NewLogger())
	// app.Use(csrf.New())

	return &Server{
		app,
		cfg,
	}
}
