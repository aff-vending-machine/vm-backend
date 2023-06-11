package store_branch

import (
	"context"

	"vm-backend/internal/layer/usecase/store_branch/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Update(ctx context.Context, req *request.Update) error {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	data := req.ToUpdate()
	_, err := uc.storeBranchRepo.Update(ctx, query, data)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to update store branch")
		return errors.Wrapf(err, "unable to update store branch %d", req.ID)
	}

	return nil
}
