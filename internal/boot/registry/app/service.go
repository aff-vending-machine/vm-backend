package app

import (
	"github.com/aff-vending-machine/vm-backend/internal/boot/registry"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/api/rpc"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/crypto/password"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/crypto/token"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository/machine"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository/machine_slot"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository/payment_channel"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository/product"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository/role"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository/transaction"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository/user"
)

func NewService(module Module) registry.Service {
	return registry.Service{
		API: registry.APIService{
			RPC: rpc.New(module.RabbitMQ.Connection),
		},
		Crypto: registry.CryptoService{
			Password: password.New(module.Config.App.BCrypt),
			Token:    token.New(module.Config.App.JWT),
		},
		Repository: registry.RepositoryService{
			Machine:        machine.New(module.PostgreSQL.DB),
			MachineSlot:    machine_slot.New(module.PostgreSQL.DB),
			PaymentChannel: payment_channel.New(module.PostgreSQL.DB),
			Product:        product.New(module.PostgreSQL.DB),
			Role:           role.New(module.PostgreSQL.DB),
			Transaction:    transaction.New(module.PostgreSQL.DB),
			User:           user.New(module.PostgreSQL.DB),
		},
	}
}
