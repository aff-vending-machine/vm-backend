package sync_rpc

import (
	"vm-backend/internal/core/domain/sync"
	"vm-backend/internal/core/infrastructure/network/rabbitmq"
)

type rpcImpl struct {
	*rabbitmq.Client
}

func NewAPI(conn *rabbitmq.Connection) sync.API {
	return &rpcImpl{
		rabbitmq.NewClient(conn),
	}
}
