package app

import (
	"github.com/aff-vending-machine/vm-backend/internal/boot/registry"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/auth"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/machine"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/machine_slot"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/payment_channel"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/product"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/report"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/role"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/sync"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/transaction"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/user"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/topic/sync_topic"
)

func NewTransport(uc registry.Usecase) registry.Transport {
	return registry.Transport{
		HTTP: registry.HTTPTransport{
			Auth:           auth.New(uc.Auth),
			Machine:        machine.New(uc.Machine),
			MachineSlot:    machine_slot.New(uc.MachineSlot),
			PaymentChannel: payment_channel.New(uc.PaymentChannel),
			Product:        product.New(uc.Product),
			Report:         report.New(uc.Report),
			Role:           role.New(uc.Role),
			Sync:           sync.New(uc.Sync),
			Transaction:    transaction.New(uc.Transaction),
			User:           user.New(uc.User),
		},
		Topic: registry.TopicTransport{
			Sync: sync_topic.New(uc.Sync),
		},
	}
}
