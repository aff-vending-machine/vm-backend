package account_role

import (
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/core/infra/strorage/postgresql/service"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
	service.Repository[account.Role]
}

func NewRepository(db *gorm.DB) account.RoleRepository {
	db.AutoMigrate(&account.Role{})
	return &repositoryImpl{
		db,
		service.New[account.Role](db),
	}
}
