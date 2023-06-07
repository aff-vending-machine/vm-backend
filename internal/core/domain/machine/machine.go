package machine

import (
	"context"
	"time"

	"vm-backend/internal/core/domain/store"
	"vm-backend/internal/core/infrastructure/strorage/postgresql/service"
	"vm-backend/internal/layer/usecase/machine/request"
	"vm-backend/internal/layer/usecase/machine/response"

	"github.com/gofiber/fiber/v2"
)

type Machine struct {
	ID                  uint          `json:"id" gorm:"primarykey"`
	BranchID            *uint         `json:"branch_id"`
	Branch              *store.Branch `json:"branch,omitempty"`
	Slots               []Slot        `json:"slots" gorm:"foreignKey:MachineID"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
	Name                string        `json:"name"`
	SerialNumber        string        `json:"serial_number" gorm:"uniqueIndex"`
	Location            string        `json:"location"`
	Type                string        `json:"type"`
	Vendor              string        `json:"vendor"`
	Status              string        `json:"status"`
	RegisterCount       int           `json:"register_count"`
	SyncTime            *time.Time    `json:"sync_time"`
	SyncSlotTime        *time.Time    `json:"sync_slot_tsime"`
	SyncChannelTime     *time.Time    `json:"sync_channel_tsime"`
	SyncTransactionTime *time.Time    `json:"sync_transaction_time"`
}

func (e Machine) TableName() string {
	return "machines"
}

type Repository interface {
	service.Repository[Machine]
}

type Usecase interface {
	Count(ctx context.Context, req *request.Filter) (int64, error)
	Create(ctx context.Context, req *request.Create) (uint, error)
	Delete(ctx context.Context, req *request.Delete) error
	Get(ctx context.Context, req *request.Get) (*response.Machine, error)
	List(ctx context.Context, req *request.Filter) ([]response.Machine, error)
	Update(ctx context.Context, req *request.Update) error
}

type Transport interface {
	Read(ctx *fiber.Ctx) error    // GET 	{machines}
	Count(ctx *fiber.Ctx) error   // GET 	{machines/count}
	ReadOne(ctx *fiber.Ctx) error // GET 	{machines/:id}
	Create(ctx *fiber.Ctx) error  // POST 	{machines}
	Update(ctx *fiber.Ctx) error  // PUT 	{machines/:id}
	Delete(ctx *fiber.Ctx) error  // DELETE {machines/:id}
}
