package product

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-backend/internal/core/module/repository"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	*repository.Template[entity.Product]
}

func New(db *gorm.DB) *repositoryImpl {
	based := repository.New[entity.Product](db)
	db.AutoMigrate(&entity.Product{})
	return &repositoryImpl{based}
}
