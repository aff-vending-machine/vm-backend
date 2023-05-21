package catalog_group

import (
	"vm-backend/internal/core/domain/catalog"
	"vm-backend/internal/core/infrastructure/strorage/postgresql/service"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
	service.Repository[catalog.Group]
}

func NewRepository(db *gorm.DB) catalog.GroupRepository {
	db.AutoMigrate(&catalog.Group{})
	return &repositoryImpl{
		db,
		service.New[catalog.Group](db),
	}
}
