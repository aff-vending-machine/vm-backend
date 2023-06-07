package sync

import (
	"context"
	"encoding/json"
	"vm-backend/internal/core/domain/sync/models"
	"vm-backend/internal/core/infra/network/rabbitmq"
	"vm-backend/internal/layer/usecase/sync/request"

	"github.com/gofiber/fiber/v2"
)

type Request[Q, D any] struct {
	Query *Q `json:"query,omitempty"`
	Data  *D `json:"data,omitempty"`
}

type Response[D any] struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    *D      `json:"data,omitempty"`
	Error   *string `json:"error,omitempty"`
}

func (r *Request[Q, D]) ToBytes() []byte {
	breq, _ := json.Marshal(r)
	return breq
}

func (r *Response[D]) IsSuccess() bool {
	return r.Code == 200
}

type API interface {
	ClearTransactions(ctx context.Context, target string, ids []uint) error
	GetChannels(ctx context.Context, target string) ([]models.Channel, error)
	GetMachine(ctx context.Context, target string) (*models.Machine, error)
	GetSlots(ctx context.Context, target string) ([]models.Slot, error)
	GetTransactions(ctx context.Context, target string) ([]models.Transaction, error)
	SetChannels(ctx context.Context, target string, data []models.Channel) error
	SetMachine(ctx context.Context, target string, data *models.Machine) error
	SetSlots(ctx context.Context, target string, data []models.Slot) error
}

type Usecase interface {
	FetchChannels(context.Context, *request.Sync) error
	FetchMachine(context.Context, *request.Sync) error
	FetchSlots(context.Context, *request.Sync) error
	PullTransactions(context.Context, *request.Sync) error
	PushChannels(context.Context, *request.Sync) error
	PushMachine(context.Context, *request.Sync) error
	PushSlots(context.Context, *request.Sync) error
	RegisterMachine(context.Context, *request.RegisterMachine) error
}

type HTTPTransport interface {
	FetchMachine(ctx *fiber.Ctx) error     // POST 	{sync/:machine_id/fetch}
	PushMachine(ctx *fiber.Ctx) error      // POST 	{sync/:machine_id/push}
	FetchChannels(ctx *fiber.Ctx) error    // POST 	{sync/:machine_id/channels/fetch}
	PushChannels(ctx *fiber.Ctx) error     // POST 	{sync/:machine_id/channels/push}
	FetchSlots(ctx *fiber.Ctx) error       // POST 	{sync/:machine_id/slots/fetch}
	PushSlots(ctx *fiber.Ctx) error        // POST 	{sync/:machine_id/slots/push}
	PullTransactions(ctx *fiber.Ctx) error // POST 	{sync/:machine_id/transactions/pull}
}

type AMQPTransport interface {
	RegisterMachine(ctx *rabbitmq.Ctx) error // {center.machine.register}
}
