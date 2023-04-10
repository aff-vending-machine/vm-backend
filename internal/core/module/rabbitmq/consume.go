package rabbitmq

import (
	"sync/atomic"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

// Channel amqp.Channel wapper
type Channel struct {
	*amqp091.Channel
	closed int32
}

// IsClosed indicate closed by developer
func (ch *Channel) IsClosed() bool {
	return (atomic.LoadInt32(&ch.closed) == 1)
}

// Close ensure closed flag set
func (ch *Channel) Close() error {
	if ch.IsClosed() {
		return amqp091.ErrClosed
	}

	atomic.StoreInt32(&ch.closed, 1)

	return ch.Channel.Close()
}

// Consume wrap amqp.Channel.Consume, the returned delivery will end only when channel closed by developer
func (ch *Channel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp091.Table) (<-chan amqp091.Delivery, error) {
	deliveries := make(chan amqp091.Delivery)

	go func() {
		for {
			time.Sleep(time.Second)
			if ch.Channel.IsClosed() {
				time.Sleep(2 * time.Second)
			}

			d, err := ch.Channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
			if err != nil {
				log.Warn().Err(err).Msg("unable to create consume")
				continue
			}

			for msg := range d {
				deliveries <- msg
			}

			time.Sleep(time.Second)
			if ch.IsClosed() {
				break
			}
		}
	}()

	return deliveries, nil
}
