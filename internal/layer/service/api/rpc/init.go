package rpc

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/rabbitmq"
	"github.com/aff-vending-machine/vm-backend/internal/core/module/rabbitmq/rpc"
)

type rpcImpl struct {
	*rpc.Client
}

func New(conn *rabbitmq.Connection) *rpcImpl {
	return &rpcImpl{
		rpc.NewClient(conn),
	}
}
