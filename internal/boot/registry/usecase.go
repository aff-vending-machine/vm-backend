package registry

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth"
	auth_usecase "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/usecase"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine"
	machine_usecase "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine/usecase"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot"
	machine_slot_usecase "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine_slot/usecase"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel"
	payment_channel_usecase "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/payment_channel/usecase"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/product"
	product_usecase "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/product/usecase"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role"
	role_usecase "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role/usecase"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/sync"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/system"
	system_usecase "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/system/usecase"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user"
	user_usecase "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/usecase"
)

// Usecase layers
type Usecase struct {
	Auth           interface{ auth.Usecase }
	Machine        interface{ machine.Usecase }
	MachineSlot    interface{ machine_slot.Usecase }
	PaymentChannel interface{ payment_channel.Usecase }
	Product        interface{ product.Usecase }
	Report         usecase.Report
	Role           interface{ role.Usecase }
	Sync           usecase.Sync
	System         interface{ system.Usecase }
	Transaction    usecase.Transaction
	User           interface{ user.Usecase }
}

func NewUsecase(adapter Service) Usecase {
	return Usecase{
		auth_usecase.New(
			adapter.Repository.Role,
			adapter.Repository.User,
			adapter.Crypto.Password,
			adapter.Crypto.Token,
		),
		machine_usecase.New(
			adapter.Repository.Machine,
			adapter.Repository.MachineSlot,
		),
		machine_slot_usecase.New(
			adapter.Repository.Machine,
			adapter.Repository.MachineSlot,
			adapter.Repository.Product,
		),
		payment_channel_usecase.New(
			adapter.Repository.PaymentChannel,
		),
		product_usecase.New(
			adapter.Repository.Product,
		),
		report.New(
			adapter.Repository.Machine,
			adapter.Repository.MachineSlot,
			adapter.Repository.Transaction,
		),
		role_usecase.New(
			adapter.Repository.Role,
		),
		sync.New(
			adapter.API.RPC,
			adapter.Repository.Machine,
			adapter.Repository.MachineSlot,
			adapter.Repository.Product,
			adapter.Repository.Transaction,
		),
		system_usecase.New(),
		transaction.New(
			adapter.API.RPC,
			adapter.Repository.Machine,
			adapter.Repository.Transaction,
		),
		user_usecase.New(
			adapter.Repository.User,
			adapter.Repository.Role,
			adapter.Crypto.Password,
		),
	}
}
