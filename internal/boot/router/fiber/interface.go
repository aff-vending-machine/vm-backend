package fiber

import (
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/core/domain/catalog"
	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/core/domain/payment"
	"vm-backend/internal/core/domain/report"
	"vm-backend/internal/core/domain/sync"
)

type Transport struct {
	Account            account.Transport
	AccountRole        account.RoleTransport
	AccountUser        account.UserTransport
	CatalogGroup       catalog.GroupTransport
	CatalogProduct     catalog.ProductTransport
	Machine            machine.Transport
	MachineSlot        machine.SlotTransport
	PaymentChannel     payment.ChannelTransport
	PaymentTransaction payment.TransactionTransport
	Report             report.Transport
	Sync               sync.HTTPTransport
}
