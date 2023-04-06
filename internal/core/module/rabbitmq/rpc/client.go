package rpc

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/rabbitmq"
)

type Client struct {
	Conn *rabbitmq.Connection
}

func NewClient(conn *rabbitmq.Connection) *Client {
	return &Client{
		Conn: conn,
	}
}
