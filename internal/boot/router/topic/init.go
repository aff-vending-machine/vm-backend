package topic

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/rabbitmq"
)

type server struct {
	*rabbitmq.Server
}

func New(app *rabbitmq.Wrapper) *server {
	return &server{
		rabbitmq.NewServer(app.Connection),
	}
}
