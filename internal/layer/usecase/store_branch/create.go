package store_branch

import (
	"context"

	"vm-backend/internal/core/domain/store"
	"vm-backend/internal/layer/usecase/store_branch/request"

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

	entity := makeStoreBranch(req)
	_, err := uc.storeBranchRepo.Create(ctx, entity)
	if err != nil {
		log.Error().Err(err).Interface("entity", entity).Msg("unable to create branch")
		return 0, errors.Wrap(err, "unable to create branch")
	}

	return entity.ID, nil
}

func makeStoreBranch(req *request.Create) *store.Branch {
	return &store.Branch{
		Name:     req.Name,
		Location: req.Location,
		IsEnable: req.IsEnable,
	}
}
