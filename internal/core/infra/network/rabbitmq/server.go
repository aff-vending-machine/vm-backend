package rabbitmq

import (
	"fmt"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Conn   *Connection
	stacks map[string]*Handler
}

func NewServer(conn *Connection) *Server {
	return &Server{
		Conn:   conn,
		stacks: make(map[string]*Handler),
	}
}

func (s *Server) Listen(queue string) error {
	if s.Conn.IsClosed() {
		return fmt.Errorf("connection closed")
	}

	channel, err := s.Conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	q, err := channel.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return err
	}

	messages, err := channel.Consume(
		q.Name,       // queue name
		"app-center", // name
		false,        // autoAck
		false,        // exclusive
		false,        // noLocal
		false,        // noWait
		nil,          // args
	)
	if err != nil {
		return err
	}

	var errs error
	for msg := range messages {
		log.Debug().Str("correlation_id", msg.CorrelationId).Str("key", msg.RoutingKey).Str("reply-to", msg.ReplyTo).Msg("rabbitmq: received")
		if msg.ReplyTo == "" {
			errs = s.topic(channel, msg)
		} else {
			errs = s.rpc(channel, msg)
		}

		if errs != nil {
			msg.Nack(false, false)
			log.Error().Err(errs).Str("correlation_id", msg.CorrelationId).Str("routing-key", msg.RoutingKey).Msg("unable to process message delivery")
		} else {
			msg.Ack(false)
		}

		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func (s *Server) Register(routingKey string, handler Handler) {
	s.stacks[routingKey] = &handler
}

func (s *Server) rpc(channel *Channel, msg amqp091.Delivery) error {
	if msg.CorrelationId == "" {
		return fmt.Errorf("rpc: no correlation ID")
	}

	routingKey := msg.Headers["routing-key"]
	if routingKey == nil {
		return fmt.Errorf("rpc: no routing Key")
	}

	key, ok := routingKey.(string)
	if !ok {
		return fmt.Errorf("rpc: routing key is not string: %t", routingKey)
	}

	handler := s.stacks[key]
	if handler == nil {
		return fmt.Errorf("rpc: no routing key registered")
	}

	ctx := NewContext(msg)
	err := (*handler)(ctx)
	if err != nil {
		ctx.InternalServer(err)
	}

	channel.PublishWithContext(
		ctx.UserContext,
		"",
		msg.ReplyTo,
		false,
		false,
		ctx.Publishing,
	)

	return nil
}

func (s *Server) topic(channel *Channel, msg amqp091.Delivery) error {
	handler := s.stacks[msg.RoutingKey]
	if handler == nil {
		return fmt.Errorf("topic: no routing key registered")
	}

	ctx := NewContext(msg)
	err := (*handler)(ctx)
	if err != nil {
		return err
	}

	return nil
}
