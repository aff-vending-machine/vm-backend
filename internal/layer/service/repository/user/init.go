package user

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-backend/internal/core/module/repository"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	*repository.Template[entity.User]
}

func New(db *gorm.DB) *repositoryImpl {
	based := repository.New[entity.User](db)
	return &repositoryImpl{based}
}
