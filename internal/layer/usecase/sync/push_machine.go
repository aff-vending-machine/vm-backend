package sync

import (
	"context"
	"strings"

	"vm-backend/internal/core/domain/sync/models"
	"vm-backend/internal/layer/usecase/sync/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) PushMachine(ctx context.Context, req *request.Sync) error {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return errors.Wrap(err, "unable to validate request")
	}

	query := req.ToMachineQuery()
	machine, err := uc.machineRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find machine")
		return errors.Wrap(err, "unable to find machine")
	}

	model := models.FromMachine(machine)
	err = uc.syncAPI.SetMachine(ctx, machine.SerialNumber, model)
	if err != nil && strings.Contains(err.Error(), "internal") {
		uc.updateMachineStatus(ctx, query, "maintenance")
		log.Error().Err(err).Str("target", machine.SerialNumber).Interface("machine", model).Msg("machine is offline")
		return errors.Wrapf(err, "real machine %s is not ready", machine.SerialNumber)
	}
	if err != nil {
		uc.updateMachineStatus(ctx, query, "offline")
		log.Error().Err(err).Str("target", machine.SerialNumber).Interface("machine", model).Msg("unable to sync machine")
		return errors.Wrapf(err, "unable to sync machine %s", machine.SerialNumber)
	}

	return nil
}
