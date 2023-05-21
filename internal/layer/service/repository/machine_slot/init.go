package machine_slot

import (
	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/core/infrastructure/strorage/postgresql/service"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
	service.Repository[machine.Slot]
}

func NewRepository(db *gorm.DB) machine.SlotRepository {
	db.AutoMigrate(&machine.Slot{})
	return &repositoryImpl{
		db,
		service.New[machine.Slot](db),
	}
}
