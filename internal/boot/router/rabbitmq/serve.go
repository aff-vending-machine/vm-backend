package rabbitmq

import (
	"github.com/rs/zerolog/log"
)

func (s *routerImpl) Serve(transport Transport) {
	s.routeSync(transport.Sync)

	go s.Listen(s.Queue)
	log.Debug().Str("queue", s.Queue).Msg("rabbitmq server listening ...")
}
