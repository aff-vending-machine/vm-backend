package catalog_group

import (
	"context"

	"vm-backend/internal/layer/usecase/catalog_group/request"
	"vm-backend/internal/layer/usecase/catalog_group/response"
	"vm-backend/pkg/helpers/conv"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Get(ctx context.Context, req *request.Get) (*response.Group, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToQuery()
	entity, err := uc.catalogGroupRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find group")
		return nil, errors.Wrapf(err, "unable to find group %d", req.ID)
	}

	return conv.StructTo[response.Group](entity)
}
