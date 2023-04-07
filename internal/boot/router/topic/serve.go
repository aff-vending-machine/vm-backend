package topic

import (
	"github.com/aff-vending-machine/vm-backend/internal/boot/registry"
	"github.com/rs/zerolog/log"
)

func (s *server) Serve(queue string, driver registry.TopicTransport) {
	s.routeSync(driver.Sync)

	go s.Listen(queue)

	log.Info().Str("queue", queue).Msg("topic server listen")
}
