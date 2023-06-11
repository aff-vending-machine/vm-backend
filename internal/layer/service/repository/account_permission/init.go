package account_permission

import (
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/core/infra/strorage/postgresql/service"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
	service.Repository[account.Permission]
}

func NewRepository(db *gorm.DB) account.PermissionRepository {
	db.AutoMigrate(&account.Permission{})
	return &repositoryImpl{
		db,
		service.New[account.Permission](db),
	}
}
