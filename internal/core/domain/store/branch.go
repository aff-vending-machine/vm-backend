package store

import (
	"context"
	"time"
	"vm-backend/internal/core/infra/strorage/postgresql/service"
	"vm-backend/internal/layer/usecase/store_branch/request"
	"vm-backend/internal/layer/usecase/store_branch/response"

	"github.com/gofiber/fiber/v2"
)

type Branch struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" gorm:"uniqueIndex"`
	Location  string    `json:"location" gorm:"uniqueIndex"`
	IsEnable  bool      `json:"is_enable"`
}

func (e Branch) TableName() string {
	return "store_branches"
}

type BranchRepository interface {
	service.Repository[Branch]
}

type BranchUsecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Get(context.Context, *request.Get) (*response.Branch, error)
	List(context.Context, *request.Filter) ([]response.Branch, error)
	Create(context.Context, *request.Create) (uint, error)
	Update(context.Context, *request.Update) error
	Delete(context.Context, *request.Delete) error
}

type BranchTransport interface {
	Read(ctx *fiber.Ctx) error    // GET		{branchs}
	Count(ctx *fiber.Ctx) error   // GET 		{branchs/count}
	ReadOne(ctx *fiber.Ctx) error // GET 		{branchs/:id}
	Create(ctx *fiber.Ctx) error  // POST		{branchs}
	Update(ctx *fiber.Ctx) error  // PUT 		{branchs/:id}
	Delete(ctx *fiber.Ctx) error  // DELETE 	{branchs/:id}
}
