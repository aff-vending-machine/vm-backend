package catalog

import (
	"context"
	"time"
	"vm-backend/internal/core/infrastructure/strorage/postgresql/service"
	"vm-backend/internal/layer/usecase/catalog_product/request"
	"vm-backend/internal/layer/usecase/catalog_product/response"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	GroupID      uint      `json:"group_id"`
	Group        *Group    `json:"group,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	SKU          string    `json:"sku" gorm:"uniqueIndex"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	ImageURL     string    `json:"image_url"`
	Barcode      string    `json:"barcode"`
	ProductPrice float64   `json:"product_price"`
	SalePrice    float64   `json:"sale_price"`
	IsEnable     bool      `json:"is_enable"`
}

func (e Product) TableName() string {
	return "catalog_products"
}

type ProductRepository interface {
	service.Repository[Product]
}

type ProductUsecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Get(context.Context, *request.Get) (*response.Product, error)
	List(context.Context, *request.Filter) ([]response.Product, error)
	Create(context.Context, *request.Create) (uint, error)
	Update(context.Context, *request.Update) error
	Delete(context.Context, *request.Delete) error
}

type ProductTransport interface {
	Read(ctx *fiber.Ctx) error    // GET 	{products}
	Count(ctx *fiber.Ctx) error   // GET 	{products/count}
	ReadOne(ctx *fiber.Ctx) error // GET 	{products/:id}
	Create(ctx *fiber.Ctx) error  // POST 	{products}
	Update(ctx *fiber.Ctx) error  // PUT 	{products/:id}
	Delete(ctx *fiber.Ctx) error  // DELETE {products/:id}
}
