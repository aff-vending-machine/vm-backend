package rpc

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/rabbitmq"
)

type rpcImpl struct {
	*rabbitmq.Client
}

func New(conn *rabbitmq.Connection) *rpcImpl {
	return &rpcImpl{
		rabbitmq.NewClient(conn),
	}
}
