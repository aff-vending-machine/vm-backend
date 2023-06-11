package catalog_product

import (
	"context"

	"vm-backend/internal/layer/usecase/catalog_product/request"

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
	_, err := uc.productRepo.Update(ctx, query, data)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Interface("data", data).Msg("unable to update catalog product")
		return errors.Wrapf(err, "unable to update catalog product %d", req.ID)
	}

	return nil
}
