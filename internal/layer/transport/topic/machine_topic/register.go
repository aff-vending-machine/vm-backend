package machine_topic

import (
	"encoding/json"

	"github.com/aff-vending-machine/vm-backend/internal/core/module/rabbitmq"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine/request"
	"github.com/aff-vending-machine/vm-backend/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (t *machineImpl) Register(c *rabbitmq.Ctx) error {
	ctx, span := trace.Start(c.UserContext)
	defer span.End()

	// request
	req, err := makeRegisterRequest(c)
	if err != nil {
		log.Error().Err(err).Msg("unable to parse request")
		trace.RecordError(span, err)
		return err
	}

	// execute usecase
	err = t.usecase.SyncRegister(ctx, req)
	if err != nil {
		log.Error().Interface("request", req).Err(err).Msg("unable to register machine")
		trace.RecordError(span, err)
		return err
	}

	return nil
}

func makeRegisterRequest(c *rabbitmq.Ctx) (*request.SyncRegister, error) {
	var req request.SyncRegister
	err := json.Unmarshal(c.Delivery.Body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
