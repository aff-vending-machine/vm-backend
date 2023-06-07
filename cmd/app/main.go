package main

import (
	"vm-backend/configs"
	"vm-backend/internal/boot/app"
	"vm-backend/pkg/boot"
	"vm-backend/pkg/log"
)

func init() {
	log.New()
}

func main() {
	// Start signal
	boot.Signal.Start()

	// Create boot with configuration
	cfg := configs.Init("env/app")
	initLog(cfg)
	cfg.Print()

	// Run main application
	app.Run(cfg)

	// Waiting for interrupt signal
	boot.Signal.Wait()
}

func initLog(cfg configs.Config) {
	log.SetOutput(log.ColorConsole())
	log.SetLogLevel(cfg.App.LogLevel)
}
