package payment_transaction

import (
	"vm-backend/internal/core/domain/payment"
	"vm-backend/internal/core/infrastructure/strorage/postgresql/service"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
	service.Repository[payment.Transaction]
}

func NewRepository(db *gorm.DB) payment.TransactionRepository {
	db.AutoMigrate(&payment.Transaction{})
	return &repositoryImpl{
		db,
		service.New[payment.Transaction](db),
	}
}
