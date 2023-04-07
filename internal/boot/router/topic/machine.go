package topic

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/topic"
)

func (s *server) routeSync(endpoint topic.Sync) {
	s.Register("center.machine.register", endpoint.Register)
}
