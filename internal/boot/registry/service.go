package registry

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/crypto"
	"github.com/aff-vending-machine/vm-backend/internal/layer/service/repository"
)

// Interface Adapter layers (driven)
type Service struct {
	API        APIService
	Crypto     CryptoService
	Repository RepositoryService
}

type APIService struct {
	RPC api.RPC
}

type CryptoService struct {
	Password crypto.Password
	Token    crypto.Token
}

type RepositoryService struct {
	Machine        repository.Machine
	MachineSlot    repository.MachineSlot
	PaymentChannel repository.PaymentChannel
	Product        repository.Product
	Role           repository.Role
	Transaction    repository.Transaction
	User           repository.User
}
