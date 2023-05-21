package app

import (
	"vm-backend/internal/layer/usecase/account"
	"vm-backend/internal/layer/usecase/account_role"
	"vm-backend/internal/layer/usecase/account_user"
	"vm-backend/internal/layer/usecase/catalog_group"
	"vm-backend/internal/layer/usecase/catalog_product"
	"vm-backend/internal/layer/usecase/machine"
	"vm-backend/internal/layer/usecase/machine_slot"
	"vm-backend/internal/layer/usecase/payment_channel"
	"vm-backend/internal/layer/usecase/payment_transaction"
	"vm-backend/internal/layer/usecase/report"
	"vm-backend/internal/layer/usecase/sync"
)

func NewUsecase(service Service) Usecase {
	return Usecase{
		Account:            account.NewUsecase(service.Repository.AccountRole, service.Repository.AccountUser, service.Management.AccountPassword, service.Management.AccountToken),
		AccountRole:        account_role.NewUsecase(service.Repository.AccountRole),
		AccountUser:        account_user.NewUsecase(service.Repository.AccountUser, service.Repository.AccountRole, service.Management.AccountPassword),
		CatalogGroup:       catalog_group.NewUsecase(service.Repository.CatalogGroup),
		CatalogProduct:     catalog_product.NewUsecase(service.Repository.CatalogProduct),
		Machine:            machine.NewUsecase(service.Repository.Machine),
		MachineSlot:        machine_slot.NewUsecase(service.Repository.Machine, service.Repository.MachineSlot, service.Repository.CatalogProduct),
		PaymentChannel:     payment_channel.NewUsecase(service.Repository.PaymentChannel),
		PaymentTransaction: payment_transaction.NewUsecase(service.Repository.Machine, service.Repository.PaymentTransaction),
		Report:             report.NewUsecase(service.Repository.Machine, service.Repository.MachineSlot, service.Repository.PaymentChannel, service.Repository.PaymentTransaction),
		Sync:               sync.NewUsecase(service.API.SyncAPI, service.Repository.PaymentChannel, service.Repository.Machine, service.Repository.MachineSlot, service.Repository.CatalogProduct, service.Repository.PaymentTransaction),
	}
}
