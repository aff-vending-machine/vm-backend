package account_user

import (
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/core/infra/strorage/postgresql/service"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
	service.Repository[account.User]
}

func NewRepository(db *gorm.DB) account.UserRepository {
	db.AutoMigrate(&account.User{})
	return &repositoryImpl{
		db,
		service.New[account.User](db),
	}
}
