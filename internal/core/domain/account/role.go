package account

import (
	"context"
	"time"
	"vm-backend/internal/core/infra/strorage/postgresql/service"
	"vm-backend/internal/layer/usecase/account_role/request"
	"vm-backend/internal/layer/usecase/account_role/response"

	"github.com/gofiber/fiber/v2"
)

type Role struct {
	ID          uint         `json:"id" gorm:"primarykey"`
	Permissions []Permission `json:"permissions"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Name        string       `json:"name" gorm:"uniqueIndex"`
}

func (e Role) TableName() string {
	return "account_roles"
}

func (e Role) HasPermission(scope string) int {
	for _, permission := range e.Permissions {
		if permission.HasScope(scope) {
			return permission.Level
		}
	}

	return 0
}

type RoleRepository interface {
	service.Repository[Role]
}

type RoleUsecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Get(context.Context, *request.Get) (*response.Role, error)
	List(context.Context, *request.Filter) ([]response.Role, error)
	Create(context.Context, *request.Create) (uint, error)
	Update(context.Context, *request.Update) error
	Delete(context.Context, *request.Delete) error
}

type RoleTransport interface {
	Read(ctx *fiber.Ctx) error    // GET 	{roles}
	Count(ctx *fiber.Ctx) error   // GET 	{roles/count}
	ReadOne(ctx *fiber.Ctx) error // GET 	{roles/:id}
	Create(ctx *fiber.Ctx) error  // POST	{roles}
	Update(ctx *fiber.Ctx) error  // PUT 	{roles/:id}
	Delete(ctx *fiber.Ctx) error  // DELETE	{roles/:id}
}
