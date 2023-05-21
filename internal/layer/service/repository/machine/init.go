package machine

import (
	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/core/infrastructure/strorage/postgresql/service"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
	service.Repository[machine.Machine]
}

func NewRepositroy(db *gorm.DB) machine.Repository {
	db.AutoMigrate(&machine.Machine{})
	return &repositoryImpl{
		db,
		service.New[machine.Machine](db),
	}
}
