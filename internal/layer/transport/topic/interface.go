package topic

import "github.com/aff-vending-machine/vm-backend/internal/core/module/rabbitmq"

type Sync interface {
	Register(ctx *rabbitmq.Ctx) error
}
