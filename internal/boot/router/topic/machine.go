package topic

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/transport/topic"
)

func (s *server) routeMachine(endpoint topic.Machine) {
	s.Register("center.machine.register", endpoint.Register)
}
