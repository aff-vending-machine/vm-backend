package machine

import (
	"context"
	"time"

	"vm-backend/internal/core/domain/catalog"
	"vm-backend/internal/core/infrastructure/strorage/postgresql/service"
	"vm-backend/internal/layer/usecase/machine_slot/request"
	"vm-backend/internal/layer/usecase/machine_slot/response"

	"github.com/gofiber/fiber/v2"
)

type Slot struct {
	ID        uint             `json:"id" gorm:"primarykey"`
	MachineID uint             `json:"machine_id"`
	Machine   *Machine         `json:"machine,omitempty"`
	ProductID uint             `json:"product_id"`
	Product   *catalog.Product `json:"product,omitempty"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	Code      string           `json:"code"`
	Stock     int              `json:"stock"`
	Capacity  int              `json:"capacity"`
	IsEnable  bool             `json:"is_enable"`
}

func (e Slot) TableName() string {
	return "machine_slots"
}

type SlotRepository interface {
	service.Repository[Slot]
}

type SlotUsecase interface {
	BulkUpdate(context.Context, *request.BulkUpdate) error
	Count(context.Context, *request.Filter) (int64, error)
	Create(context.Context, *request.Create) (uint, error)
	Delete(context.Context, *request.Delete) error
	Get(context.Context, *request.Get) (*response.Slot, error)
	List(context.Context, *request.Filter) ([]response.Slot, error)
	Update(context.Context, *request.Update) error
}

type SlotTransport interface {
	Read(ctx *fiber.Ctx) error       // GET 	{machines/:machine_id/slots}
	Count(ctx *fiber.Ctx) error      // GET 	{machines/:machine_id/slots/count}
	ReadOne(ctx *fiber.Ctx) error    // GET 	{machines/:machine_id/slots/:id}
	Create(ctx *fiber.Ctx) error     // POST 	{machines/:machine_id/slots}
	BulkUpdate(ctx *fiber.Ctx) error // PUT 	{machines/:machine_id/slots/bulk}
	Update(ctx *fiber.Ctx) error     // PUT 	{machines/:machine_id/slots/:id}
	Delete(ctx *fiber.Ctx) error     // DELETE 	{machines/:machine_id/slots/:id}
}
