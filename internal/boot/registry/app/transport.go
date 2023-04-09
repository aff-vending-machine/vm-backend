package app

import (
	"github.com/aff-vending-machine/vm-backend/internal/boot/registry"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/auth_http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/machine_http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/machine_slot_http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/payment_channel_http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/product_http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/report_http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/role_http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/sync_http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/transaction_http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http/user_http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/topic/sync_topic"
)

func NewTransport(uc registry.Usecase) registry.Transport {
	return registry.Transport{
		HTTP: registry.HTTPTransport{
			Auth:           auth_http.New(uc.Auth),
			Machine:        machine_http.New(uc.Machine),
			MachineSlot:    machine_slot_http.New(uc.MachineSlot),
			PaymentChannel: payment_channel_http.New(uc.PaymentChannel),
			Product:        product_http.New(uc.Product),
			Report:         report_http.New(uc.Report),
			Role:           role_http.New(uc.Role),
			Sync:           sync_http.New(uc.Sync),
			Transaction:    transaction_http.New(uc.Transaction),
			User:           user_http.New(uc.User),
		},
		Topic: registry.TopicTransport{
			Sync: sync_topic.New(uc.Sync),
		},
	}
}
