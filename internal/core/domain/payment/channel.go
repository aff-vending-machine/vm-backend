package payment

import (
	"context"
	"time"
	"vm-backend/internal/core/infrastructure/strorage/postgresql/service"
	"vm-backend/internal/layer/usecase/payment_channel/request"
	"vm-backend/internal/layer/usecase/payment_channel/response"

	"github.com/gofiber/fiber/v2"
)

type Channel struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	MachineID    uint      `json:"machine_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	Channel      string    `json:"channel"`
	Vendor       string    `json:"vendor"`
	IsEnable     bool      `json:"is_enable"`
	Host         string    `json:"host"`
	MerchantID   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	BillerCode   string    `json:"biller_code"`
	BillerID     string    `json:"biller_id"`
	Token        string    `json:"token"`
	StoreID      string    `json:"store_id"`
	TerminalID   string    `json:"terminal_id"`
}

func (c *Channel) TableName() string {
	return "payment_channels"
}

type ChannelRepository interface {
	service.Repository[Channel]
}

type ChannelUsecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Get(context.Context, *request.Get) (*response.Channel, error)
	List(context.Context, *request.Filter) ([]response.Channel, error)
	Create(context.Context, *request.Create) (uint, error)
	Update(context.Context, *request.Update) error
	Delete(context.Context, *request.Delete) error
	Enable(context.Context, *request.Enable) error
	Disable(context.Context, *request.Disable) error
}

type ChannelTransport interface {
	Read(ctx *fiber.Ctx) error    // GET	{channels}
	Count(ctx *fiber.Ctx) error   // GET 	{channels/count}
	ReadOne(ctx *fiber.Ctx) error // GET 	{channels/:id}
	Create(ctx *fiber.Ctx) error  // POST	{channels}
	Enable(ctx *fiber.Ctx) error  // POST 	{channels/:id/enable}
	Disable(ctx *fiber.Ctx) error // POST 	{channels/:id/disable}
	Update(ctx *fiber.Ctx) error  // PUT 	{channels/:id}
	Delete(ctx *fiber.Ctx) error  // DELETE {channels/:id}
}
