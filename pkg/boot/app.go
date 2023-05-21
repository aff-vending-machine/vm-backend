package boot

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

var Signal = NewApp()

type App struct {
	Context      context.Context
	Cancel       context.CancelFunc
	stopChan     chan bool
	shutdownChan chan struct{}
}

func NewApp() *App {
	ctx, cancel := context.WithCancel(context.TODO())
	return &App{
		Context:      ctx,
		Cancel:       cancel,
		stopChan:     make(chan bool, 1),
		shutdownChan: make(chan struct{}),
	}
}

func (a *App) Start() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-signals:
		case <-a.stopChan:
		}

		a.shutdown()
	}()

}

func (a *App) Stop() {
	a.stopChan <- true
	time.Sleep(3 * time.Second)
}

func (a *App) Wait() {
	<-a.shutdownChan
}

func (a *App) shutdown() {
	a.Cancel()
	if err := cleanup(); err != nil {
		log.Error().Err(err).Msg("Error during shutdown")
	}
	log.Info().Msg("Shutting down gracefully...")

	// Signal that the shutdown is complete
	close(a.shutdownChan)
	os.Exit(0)
}

func cleanup() error {
	return nil
}
