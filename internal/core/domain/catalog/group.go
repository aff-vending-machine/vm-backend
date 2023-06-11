package catalog

import (
	"context"
	"time"
	"vm-backend/internal/core/infra/strorage/postgresql/service"
	"vm-backend/internal/layer/usecase/catalog_group/request"
	"vm-backend/internal/layer/usecase/catalog_group/response"

	"github.com/gofiber/fiber/v2"
)

type Group struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	Products    []Product `json:"products" gorm:"foreignKey:GroupID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsEnable    bool      `json:"is_enable"`
}

func (e Group) TableName() string {
	return "catalog_groups"
}

type GroupRepository interface {
	service.Repository[Group]
}

type GroupUsecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Get(context.Context, *request.Get) (*response.Group, error)
	List(context.Context, *request.Filter) ([]response.Group, error)
	Create(context.Context, *request.Create) (uint, error)
	Update(context.Context, *request.Update) error
	Delete(context.Context, *request.Delete) error
}

type GroupTransport interface {
	Read(ctx *fiber.Ctx) error    // GET 	{groups}
	Count(ctx *fiber.Ctx) error   // GET 	{groups/count}
	ReadOne(ctx *fiber.Ctx) error // GET 	{groups/:id}
	Create(ctx *fiber.Ctx) error  // POST 	{groups}
	Update(ctx *fiber.Ctx) error  // PUT 	{groups/:id}
	Delete(ctx *fiber.Ctx) error  // DELETE {groups/:id}
}
