package store_branch

import (
	"vm-backend/internal/core/domain/store"
	"vm-backend/internal/core/infra/strorage/postgresql/service"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
	service.Repository[store.Branch]
}

func NewRepository(db *gorm.DB) store.BranchRepository {
	db.AutoMigrate(&store.Branch{})
	return &repositoryImpl{
		db,
		service.New[store.Branch](db),
	}
}
