package payment_channel

import (
	"vm-backend/internal/core/domain/payment"
	"vm-backend/internal/core/infra/strorage/postgresql/service"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
	service.Repository[payment.Channel]
}

func NewRepository(db *gorm.DB) payment.ChannelRepository {
	db.AutoMigrate(&payment.Channel{})
	return &repositoryImpl{
		db,
		service.New[payment.Channel](db),
	}
}
