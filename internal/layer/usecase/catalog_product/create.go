package catalog_product

import (
	"context"

	"vm-backend/internal/core/domain/catalog"
	"vm-backend/internal/layer/usecase/catalog_product/request"

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

	entity := makeCatalogProduct(req)
	_, err := uc.productRepo.Create(ctx, entity)
	if err != nil {
		log.Error().Err(err).Interface("entity", entity).Msg("unable to create product")
		return 0, errors.Wrap(err, "unable to create product")
	}

	return entity.ID, nil
}

func makeCatalogProduct(req *request.Create) *catalog.Product {
	return &catalog.Product{
		GroupID:      req.GroupID,
		SKU:          req.SKU,
		Name:         req.Name,
		Description:  req.Description,
		ImageURL:     req.ImageURL,
		Barcode:      req.Barcode,
		ProductPrice: req.ProductPrice,
		SalePrice:    req.SalePrice,
		IsEnable:     req.IsEnable,
	}
}
