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
	conf := configs.Init("env/app")
	initLog(conf)

	conf.App.Version = "2.0.0"
	conf.Print()

	// Run main application
	app.Run(conf)

	// Waiting for interrupt signal
	boot.Signal.Wait()
}

func initLog(conf configs.Config) {
	if conf.App.ENV == "local" {
		log.SetOutput(log.ColorConsole())
	}

	log.SetLogLevel(conf.App.LogLevel)
}
