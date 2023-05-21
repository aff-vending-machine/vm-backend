package rabbitmq

import "vm-backend/internal/core/domain/sync"

type Transport struct {
	Sync sync.AMQPTransport
}
