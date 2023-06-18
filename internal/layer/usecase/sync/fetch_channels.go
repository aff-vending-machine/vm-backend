package sync

import (
	"context"
	"time"

	"vm-backend/internal/layer/usecase/sync/request"
	"vm-backend/pkg/helpers/db"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) FetchChannels(ctx context.Context, req *request.Sync) error {
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

	channels, err := uc.syncAPI.GetChannels(ctx, machine.SerialNumber)
	if err != nil {
		log.Error().Err(err).Str("target", machine.SerialNumber).Msg("unable to fetch channels")
		return errors.Wrap(err, "unable to fetch channels")
	}

	for _, channel := range channels {
		query := db.NewQuery().Where("channel = ? AND machine_id = ?", channel.Name, machine.ID)
		update := channel.ToUpdate()
		uc.slotRepo.Update(ctx, query, update)
	}

	query = db.NewQuery().Where("id = ?", req.MachineID)
	update := map[string]interface{}{"sync_channel_time": time.Now()}
	_, err = uc.machineRepo.Update(ctx, query, update)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Interface("update", update).Msg("unable to update machine")
		return errors.Wrap(err, "unable to update machine")
	}

	return nil
}
