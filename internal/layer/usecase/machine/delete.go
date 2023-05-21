package machine

import (
	"context"
	"vm-backend/internal/layer/usecase/machine/request"

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
	_, err := uc.machineRepo.Delete(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to delete machine")
		return errors.Wrapf(err, "unable to delete machine %d", req.ID)
	}

	return nil
}
