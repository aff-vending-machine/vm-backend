package store_branch

import (
	"context"

	"vm-backend/internal/layer/usecase/store_branch/request"
	"vm-backend/internal/layer/usecase/store_branch/response"
	"vm-backend/pkg/helpers/conv"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) List(ctx context.Context, req *request.Filter) ([]response.Branch, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	entities, err := uc.storeBranchRepo.FindMany(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find store branchs")
		return nil, errors.Wrap(err, "unable to find store branchs")
	}

	return conv.ToArray[response.Branch](entities)
}
