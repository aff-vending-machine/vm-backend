package registry

import "github.com/aff-vending-machine/vm-backend/internal/layer/transport/http"

// Interface Adapter layers (driver)
type Transport struct {
	HTTP HTTPTransport
}

type HTTPTransport struct {
	Auth           http.Auth
	Machine        http.Machine
	MachineSlot    http.MachineSlot
	PaymentChannel http.PaymentChannel
	Product        http.Product
	Role           http.Role
	Transaction    http.Transaction
	User           http.User
}
