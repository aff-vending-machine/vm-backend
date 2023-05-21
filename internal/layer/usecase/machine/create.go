package machine

import (
	"context"
	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/layer/usecase/machine/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Create(ctx context.Context, req *request.Create) (uint, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return 0, errors.Wrap(err, "unable to validate request")
	}

	entity := makeMachine(req)
	_, err := uc.machineRepo.Create(ctx, entity)
	if err != nil {
		log.Error().Err(err).Interface("entity", entity).Msg("unable to create machine")
		return 0, errors.Wrap(err, "unable to create machine")
	}

	return entity.ID, nil
}

func makeMachine(req *request.Create) *machine.Machine {
	return &machine.Machine{
		BranchID:      req.BranchID,
		Name:          req.Name,
		SerialNumber:  req.SerialNumber,
		Location:      req.Location,
		Type:          req.Type,
		Vendor:        req.Vendor,
		Status:        "enable",
		RegisterCount: 0,
	}
}
