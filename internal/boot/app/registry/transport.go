package registry

import (
	"vm-backend/internal/boot/modules"
	"vm-backend/internal/boot/router/fiber"
	"vm-backend/internal/boot/router/rabbitmq"
	"vm-backend/internal/layer/transport/fiber/account"
	"vm-backend/internal/layer/transport/fiber/account_role"
	"vm-backend/internal/layer/transport/fiber/account_user"
	"vm-backend/internal/layer/transport/fiber/catalog_group"
	"vm-backend/internal/layer/transport/fiber/catalog_product"
	"vm-backend/internal/layer/transport/fiber/machine"
	"vm-backend/internal/layer/transport/fiber/machine_slot"
	"vm-backend/internal/layer/transport/fiber/payment_channel"
	"vm-backend/internal/layer/transport/fiber/payment_transaction"
	"vm-backend/internal/layer/transport/fiber/report"
	"vm-backend/internal/layer/transport/fiber/store_branch"
	"vm-backend/internal/layer/transport/fiber/sync_http"
	"vm-backend/internal/layer/transport/rabbitmq/sync_amqp"
)

func NewTransport(usecase modules.Usecase) modules.Transport {
	return modules.Transport{
		Fiber: fiber.Transport{
			Account:            account.NewTransport(usecase.Account),
			AccountRole:        account_role.NewTransport(usecase.AccountRole),
			AccountUser:        account_user.NewTransport(usecase.AccountUser),
			CatalogGroup:       catalog_group.NewTransport(usecase.CatalogGroup),
			CatalogProduct:     catalog_product.NewTransport(usecase.CatalogProduct),
			Machine:            machine.NewTransport(usecase.Machine),
			MachineSlot:        machine_slot.NewTransport(usecase.MachineSlot),
			PaymentChannel:     payment_channel.NewTransport(usecase.PaymentChannel),
			PaymentTransaction: payment_transaction.NewTransport(usecase.PaymentTransaction),
			Report:             report.NewTransport(usecase.Report),
			StoreBranch:        store_branch.NewTransport(usecase.StoreBranch),
			Sync:               sync_http.NewTransport(usecase.Sync),
		},
		RabbitMQ: rabbitmq.Transport{
			Sync: sync_amqp.NewTransport(usecase.Sync),
		},
	}
}
