package transaction

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-backend/internal/core/module/repository"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	*repository.Template[entity.Transaction]
}

func New(db *gorm.DB) *repositoryImpl {
	based := repository.New[entity.Transaction](db)
	db.AutoMigrate(&entity.Transaction{})
	return &repositoryImpl{based}
}