package catalog_product

import (
	"vm-backend/internal/core/domain/catalog"
	"vm-backend/internal/core/infra/strorage/postgresql/service"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
	service.Repository[catalog.Product]
}

func NewRepository(db *gorm.DB) catalog.ProductRepository {
	db.AutoMigrate(&catalog.Product{})
	return &repositoryImpl{
		db,
		service.New[catalog.Product](db),
	}
}
