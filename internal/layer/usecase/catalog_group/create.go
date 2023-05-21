package catalog_group

import (
	"context"

	"vm-backend/internal/core/domain/catalog"
	"vm-backend/internal/layer/usecase/catalog_group/request"

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

	entity := makeCatalogGroup(req)
	_, err := uc.catalogGroupRepo.Create(ctx, entity)
	if err != nil {
		log.Error().Err(err).Interface("entity", entity).Msg("unable to create group")
		return 0, errors.Wrap(err, "unable to create group")
	}

	return entity.ID, nil
}

func makeCatalogGroup(req *request.Create) *catalog.Group {
	return &catalog.Group{
		Name:        req.Name,
		Description: req.Description,
		IsEnable:    req.IsEnable,
	}
}
