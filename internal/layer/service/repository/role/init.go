package role

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-backend/internal/core/module/repository"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	*repository.Template[entity.Role]
}

func New(db *gorm.DB) *repositoryImpl {
	based := repository.New[entity.Role](db)
	db.AutoMigrate(&entity.Role{}, &entity.Permission{})
	return &repositoryImpl{based}
}
