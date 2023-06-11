package registry

import (
	"vm-backend/internal/boot/modules"
	"vm-backend/internal/layer/service/api/sync_rpc"
	"vm-backend/internal/layer/service/management/account_password"
	"vm-backend/internal/layer/service/management/account_token"
	"vm-backend/internal/layer/service/repository/account_role"
	"vm-backend/internal/layer/service/repository/account_user"
	"vm-backend/internal/layer/service/repository/catalog_group"
	"vm-backend/internal/layer/service/repository/catalog_product"
	"vm-backend/internal/layer/service/repository/machine"
	"vm-backend/internal/layer/service/repository/machine_slot"
	"vm-backend/internal/layer/service/repository/payment_channel"
	"vm-backend/internal/layer/service/repository/payment_transaction"
)

func NewService(infra modules.Infrastructure) modules.Service {
	return modules.Service{
		API: modules.APIService{
			SyncAPI: sync_rpc.NewAPI(infra.RabbitMQ.Connection),
		},
		Management: modules.ManagementService{
			AccountPassword: account_password.NewManagement(infra.App.BCrypt),
			AccountToken:    account_token.NewManagement(infra.App.JWT),
		},
		Repository: modules.RepositoryService{
			AccountRole:        account_role.NewRepository(infra.PostgreSQL.DB),
			AccountUser:        account_user.NewRepository(infra.PostgreSQL.DB),
			CatalogGroup:       catalog_group.NewRepository(infra.PostgreSQL.DB),
			CatalogProduct:     catalog_product.NewRepository(infra.PostgreSQL.DB),
			Machine:            machine.NewRepositroy(infra.PostgreSQL.DB),
			MachineSlot:        machine_slot.NewRepository(infra.PostgreSQL.DB),
			PaymentChannel:     payment_channel.NewRepository(infra.PostgreSQL.DB),
			PaymentTransaction: payment_transaction.NewRepository(infra.PostgreSQL.DB),
		},
	}
}
