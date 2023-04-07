package registry

import (
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
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role"
	role_usecase "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role/usecase"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/system"
	system_usecase "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/system/usecase"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction"
	transaction_usecase "github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/usecase"
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
	Role           interface{ role.Usecase }
	System         interface{ system.Usecase }
	Transaction    interface{ transaction.Usecase }
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
			adapter.API.RPC,
			adapter.Repository.Machine,
			adapter.Repository.MachineSlot,
		),
		machine_slot_usecase.New(
			adapter.API.RPC,
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
		role_usecase.New(
			adapter.Repository.Role,
		),
		system_usecase.New(),
		transaction_usecase.New(
			adapter.Repository.Transaction,
		),
		user_usecase.New(
			adapter.Repository.User,
			adapter.Repository.Role,
			adapter.Crypto.Password,
		),
	}
}
