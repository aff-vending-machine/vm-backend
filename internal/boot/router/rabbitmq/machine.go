package rabbitmq

import "vm-backend/internal/core/domain/sync"

func (r *routerImpl) routeSync(endpoint sync.AMQPTransport) {
	r.Register("center.machine.register", endpoint.RegisterMachine)
}
