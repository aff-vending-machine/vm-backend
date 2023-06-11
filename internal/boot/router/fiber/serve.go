package fiber

import (
	"fmt"
	"vm-backend/internal/boot/router/fiber/apiv2"

	"github.com/rs/zerolog/log"
)

func (s *routerImpl) Serve(transport Transport) {
	// root resource
	v2 := s.Group("/api/v2")
	apiv2.RouteAccount(v2, transport.Account)

	v2.Use(transport.Account.AuthorizationRequired, transport.Account.PermissionRequired)

	apiv2.RouteAccountRole(v2, transport.AccountRole)
	apiv2.RouteAccountUser(v2, transport.AccountUser)
	apiv2.RouteCatalogGroup(v2, transport.CatalogGroup)
	apiv2.RouteCatalogProduct(v2, transport.CatalogProduct)
	apiv2.RouteMachine(v2, transport.Machine)
	apiv2.RouteMachineSlot(v2, transport.MachineSlot)
	apiv2.RoutePaymentChannel(v2, transport.PaymentChannel)
	apiv2.RoutePaymentTransaction(v2, transport.PaymentTransaction)
	apiv2.RouteReport(v2, transport.Report)
	apiv2.RouteStoreBranch(v2, transport.StoreBranch)
	apiv2.RouteSync(v2, transport.Sync)

	addr := fmt.Sprintf(":%d", s.Port)
	go s.Listen(addr)
	log.Debug().Int("port", s.Port).Msg("http server listening ...")
}
