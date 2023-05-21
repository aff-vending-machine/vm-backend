package machine_slot

import (
	"context"

	"vm-backend/internal/layer/usecase/machine_slot/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) BulkUpdate(ctx context.Context, req *request.BulkUpdate) error {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return errors.Wrap(err, "unable to validate request")
	}

	if isExist := uc.isMachineExist(ctx, req.MachineID); !isExist {
		return errors.Errorf("machine %s not found", req.MachineID)
	}

	for i, data := range req.Data {
		query := req.ToQuery(data.ID)
		update := req.ToUpdate(i)
		_, err := uc.machineSlotRepo.Update(ctx, query, update)
		if err != nil {
			log.Error().Err(err).
				Int("index", i).
				Uint("id", data.ID).
				Interface("query", query).
				Interface("data", update).
				Msg("unable to update machine slot")
			continue
		}
	}

	return nil
}
