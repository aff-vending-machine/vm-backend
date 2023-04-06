package machine_slot

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-backend/internal/core/module/repository"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	*repository.Template[entity.MachineSlot]
}

func New(db *gorm.DB) *repositoryImpl {
	based := repository.New[entity.MachineSlot](db)
	db.AutoMigrate(&entity.MachineSlot{})
	return &repositoryImpl{based}
}
