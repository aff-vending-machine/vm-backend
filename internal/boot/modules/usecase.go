package modules

import (
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/core/domain/catalog"
	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/core/domain/payment"
	"vm-backend/internal/core/domain/report"
	"vm-backend/internal/core/domain/store"
	"vm-backend/internal/core/domain/sync"
)

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
	StoreBranch        store.BranchUsecase
	Sync               sync.Usecase
}
