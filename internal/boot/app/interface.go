package app

import (
	"vm-backend/internal/boot/router/fiber"
	"vm-backend/internal/boot/router/rabbitmq"
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/core/domain/catalog"
	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/core/domain/payment"
	"vm-backend/internal/core/domain/report"
	"vm-backend/internal/core/domain/sync"
)

// Interface Adapter layers (driven)
type Service struct {
	API        APIService
	Management ManagementService
	Repository RepositoryService
}

type APIService struct {
	SyncAPI sync.API
}

type ManagementService struct {
	AccountPassword account.PasswordManagement
	AccountToken    account.TokenManagement
}

type RepositoryService struct {
	AccountPermission  account.PermissionRepository
	AccountRole        account.RoleRepository
	AccountUser        account.UserRepository
	CatalogGroup       catalog.GroupRepository
	CatalogProduct     catalog.ProductRepository
	Machine            machine.Repository
	MachineSlot        machine.SlotRepository
	PaymentChannel     payment.ChannelRepository
	PaymentTransaction payment.TransactionRepository
}

// Usecase layers
type Usecase struct {
	Account            account.Usecase
	AccountRole        account.RoleUsecase
	AccountUser        account.UserUsecase
	CatalogGroup       catalog.GroupUsecase
	CatalogProduct     catalog.ProductUsecase
	Machine            machine.Usecase
	MachineSlot        machine.SlotUsecase
	PaymentChannel     payment.ChannelUsecase
	PaymentTransaction payment.TransactionUsecase
	Report             report.Usecase
	Sync               sync.Usecase
}

// Interface Adapter layers (driver)
type Transport struct {
	Fiber    fiber.Transport
	RabbitMQ rabbitmq.Transport
}
