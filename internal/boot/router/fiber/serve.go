package fiber

import (
	"github.com/aff-vending-machine/vm-backend/internal/boot/registry"
	"github.com/rs/zerolog/log"
)

func (s *server) Serve(driver registry.HTTPTransport) {
	// root resource
	v1 := s.App.Group("/api/v1")
	routeAuth(v1, driver.Auth)

	// v1.Use(driver.Auth.AuthorizationRequired, driver.Auth.PermissionRequired)

	routeMachine(v1, driver.Machine)
	routeMachineSlot(v1, driver.MachineSlot)
	routePaymentChannel(v1, driver.PaymentChannel)
	routeProduct(v1, driver.Product)
	routeRole(v1, driver.Role)
	routeSync(v1, driver.Sync)
	routeTransaction(v1, driver.Transaction)
	routeUser(v1, driver.User)

	go s.App.Listen(s.Address)

	log.Info().Str("address", s.Address).Msg("http server listen")
}
