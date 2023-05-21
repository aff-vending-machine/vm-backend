package machine_slot

import (
	"context"

	"vm-backend/internal/layer/usecase/machine_slot/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Delete(ctx context.Context, req *request.Delete) error {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	_, err := uc.machineSlotRepo.Delete(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to delete slot")
		return errors.Wrapf(err, "unable to delete slot %d at machine %d", req.ID, req.MachineID)
	}

	return nil
}
