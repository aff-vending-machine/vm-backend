package boot

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

type ClosedFn func() error

var Signal = NewApp()

type App struct {
	Context      context.Context
	Cancel       context.CancelFunc
	stopChan     chan bool
	shutdownChan chan struct{}
	ClosedFn     []ClosedFn
}

func NewApp() *App {
	ctx, cancel := context.WithCancel(context.TODO())
	return &App{
		Context:      ctx,
		Cancel:       cancel,
		stopChan:     make(chan bool, 1),
		shutdownChan: make(chan struct{}),
		ClosedFn:     make([]ClosedFn, 0),
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

func (a *App) AddClosedFn(fn ClosedFn) {
	a.ClosedFn = append(a.ClosedFn, fn)
}

func (a *App) shutdown() {
	a.Cancel()
	for _, close := range a.ClosedFn {
		if err := close(); err != nil {
			log.Error().Err(err).Msg("Error during close module")
		}
	}

	log.Info().Msg("Shutting down gracefully...")

	// Signal that the shutdown is complete
	close(a.shutdownChan)
	os.Exit(0)
}
