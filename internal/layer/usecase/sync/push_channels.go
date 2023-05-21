package sync

import (
	"context"
	"time"

	"vm-backend/internal/core/domain/sync/models"
	"vm-backend/internal/layer/usecase/sync/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) PushChannels(ctx context.Context, req *request.Sync) error {
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

	query = req.ToChannelQuery()
	channels, err := uc.channelRepo.FindMany(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find channels")
		return errors.Wrap(err, "unable to find channels")
	}

	data := models.FromChannelList(channels, machine.ID)
	err = uc.syncAPI.SetChannels(ctx, machine.SerialNumber, data)
	if err != nil {
		log.Error().Err(err).Str("target", machine.SerialNumber).Interface("data", data).Msg("unable to sync machine")
		return errors.Wrapf(err, "unable to sync machine %s", machine.SerialNumber)
	}

	query = req.ToMachineQuery()
	update := map[string]interface{}{"sync_channel_time": time.Now()}
	_, err = uc.machineRepo.Update(ctx, query, update)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Interface("update", update).Msg("unable to update machine")
		return errors.Wrapf(err, "unable to update machine %s", machine.SerialNumber)
	}

	return nil
}
