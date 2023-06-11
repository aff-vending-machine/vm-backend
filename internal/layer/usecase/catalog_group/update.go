package catalog_group

import (
	"context"

	"vm-backend/internal/layer/usecase/catalog_group/request"

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
	_, err := uc.catalogGroupRepo.Update(ctx, query, data)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to update catalog group")
		return errors.Wrapf(err, "unable to update catalog group %d", req.ID)
	}

	return nil
}
