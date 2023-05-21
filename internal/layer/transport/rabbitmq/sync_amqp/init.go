package sync_amqp

import "vm-backend/internal/core/domain/sync"

type transportImpl struct {
	usecase sync.Usecase
}

func NewTransport(uc sync.Usecase) sync.AMQPTransport {
	return &transportImpl{uc}
}
