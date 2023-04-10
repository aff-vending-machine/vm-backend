package registry

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/topic"
)

// Interface Adapter layers (driver)
type Transport struct {
	HTTP  HTTPTransport
	Topic TopicTransport
}

type HTTPTransport struct {
	Auth           http.Auth
	Machine        http.Machine
	MachineSlot    http.MachineSlot
	PaymentChannel http.PaymentChannel
	Product        http.Product
	Report         http.Report
	Role           http.Role
	Sync           http.Sync
	Transaction    http.Transaction
	User           http.User
}

type TopicTransport struct {
	Sync topic.Sync
}
