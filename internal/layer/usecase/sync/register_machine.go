package sync

import (
	"context"
	"time"

	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/layer/usecase/sync/request"
	"vm-backend/pkg/errs"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) RegisterMachine(ctx context.Context, req *request.RegisterMachine) error {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	machine, err := uc.machineRepo.FindOne(ctx, query)
	if errs.Is(err, errs.ErrNotFound) {
		entity := makeMachine(req, 1)
		_, err := uc.machineRepo.Create(ctx, entity)
		if err != nil {
			log.Error().Err(err).Interface("entity", entity).Msg("unable to create machine")
			return errors.Wrap(err, "unable to create machine")
		}
		return nil
	}
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find machine")
		return errors.Wrap(err, "unable to find machine")
	}

	update := req.ToUpdate(machine.RegisterCount)
	_, err = uc.machineRepo.Update(ctx, query, update)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Interface("update", update).Msg("unable to update machine")
		return errors.Wrap(err, "unable to update machine")
	}

	return nil
}

func makeMachine(req *request.RegisterMachine, branchID uint) *machine.Machine {
	var id *uint
	t := time.Now()
	if branchID > 0 {
		id = &branchID
	}

	return &machine.Machine{
		Name:         req.Data.Name,
		BranchID:     id,
		SerialNumber: req.Data.SerialNumber,
		Location:     req.Data.Location,
		Type:         "<auto register>",
		Vendor:       req.Data.Vendor,
		SyncTime:     &t,
		Status:       "online",
	}
}
