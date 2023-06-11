package sync_amqp

import (
	"encoding/json"

	"vm-backend/internal/core/infra/network/rabbitmq"
	"vm-backend/internal/layer/usecase/sync/request"
)

func (t *transportImpl) RegisterMachine(c *rabbitmq.Ctx) error {
	ctx := c.UserContext
	// request
	req, err := makeRegisterRequest(c)
	if err != nil {
		return err
	}

	// execute usecase
	err = t.usecase.RegisterMachine(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func makeRegisterRequest(c *rabbitmq.Ctx) (*request.RegisterMachine, error) {
	var req request.RegisterMachine
	err := json.Unmarshal(c.Delivery.Body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
