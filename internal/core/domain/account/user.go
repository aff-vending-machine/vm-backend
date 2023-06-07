package account

import (
	"context"
	"time"
	"vm-backend/internal/core/domain/store"
	"vm-backend/internal/core/infrastructure/strorage/postgresql/service"
	"vm-backend/internal/layer/usecase/account_user/request"
	"vm-backend/internal/layer/usecase/account_user/response"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint          `json:"id" gorm:"primarykey"`
	BranchID  *uint         `json:"branch_id"`
	Branch    *store.Branch `json:"branch"`
	RoleID    uint          `json:"role_id"`
	Role      Role          `json:"role"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Username  string        `json:"username" gorm:"uniqueIndex"`
	Password  string        `json:"-"`
	Contact   string        `json:"contact"`
	CreatedBy string        `json:"created_by"`
	LastLogin *time.Time    `json:"last_login"`
	LastToken string        `json:"-"`
}

func (e User) TableName() string {
	return "account_users"
}

func (e User) HasRole(name string) bool {
	return e.Role.Name == name
}

type UserRepository interface {
	service.Repository[User]
}

type UserUsecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Get(context.Context, *request.Get) (*response.User, error)
	List(context.Context, *request.Filter) ([]response.User, error)
	Create(context.Context, *request.Create) (uint, error)
	ChangeRole(context.Context, *request.ChangeRole) error
	ChangePassword(context.Context, *request.ChangePassword) error
	ResetPassword(context.Context, *request.ResetPassword) error
	Delete(context.Context, *request.Delete) error
}

type UserTransport interface {
	Read(ctx *fiber.Ctx) error           // GET		{users}
	Count(ctx *fiber.Ctx) error          // GET 	{users/count}
	ReadOne(ctx *fiber.Ctx) error        // GET 	{users/:id}
	Create(ctx *fiber.Ctx) error         // POST	{users}
	ChangeRole(ctx *fiber.Ctx) error     // POST 	{users/:id/change-role}
	ChangePassword(ctx *fiber.Ctx) error // POST 	{users/me/change-password}
	ResetPassword(ctx *fiber.Ctx) error  // POST 	{users/:id/reset-password}
	Delete(ctx *fiber.Ctx) error         // DELETE 	{users/:id}
}
