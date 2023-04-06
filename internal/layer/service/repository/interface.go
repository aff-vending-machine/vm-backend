package repository

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Machine interface {
	Count(ctx context.Context, filter []string) (int64, error)
	FindOne(ctx context.Context, filter []string) (*entity.Machine, error)
	FindMany(ctx context.Context, filter []string) ([]entity.Machine, error)
	InsertOne(ctx context.Context, ent *entity.Machine) error
	UpdateMany(ctx context.Context, filter []string, ent map[string]interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter []string) (int64, error)
}

type PaymentChannel interface {
	Count(ctx context.Context, filter []string) (int64, error)
	FindOne(ctx context.Context, filter []string) (*entity.PaymentChannel, error)
	FindMany(ctx context.Context, filter []string) ([]entity.PaymentChannel, error)
	InsertOne(ctx context.Context, ent *entity.PaymentChannel) error
	UpdateMany(ctx context.Context, filter []string, ent map[string]interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter []string) (int64, error)
}

type MachineSlot interface {
	Count(ctx context.Context, filter []string) (int64, error)
	FindOne(ctx context.Context, filter []string) (*entity.MachineSlot, error)
	FindMany(ctx context.Context, filter []string) ([]entity.MachineSlot, error)
	InsertOne(ctx context.Context, ent *entity.MachineSlot) error
	UpdateMany(ctx context.Context, filter []string, ent map[string]interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter []string) (int64, error)
}

type Product interface {
	Count(ctx context.Context, filter []string) (int64, error)
	FindOne(ctx context.Context, filter []string) (*entity.Product, error)
	FindMany(ctx context.Context, filter []string) ([]entity.Product, error)
	InsertOne(ctx context.Context, ent *entity.Product) error
	UpdateMany(ctx context.Context, filter []string, ent map[string]interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter []string) (int64, error)
}

type Role interface {
	Count(ctx context.Context, filter []string) (int64, error)
	FindOne(ctx context.Context, filter []string) (*entity.Role, error)
	FindMany(ctx context.Context, filter []string) ([]entity.Role, error)
	InsertOne(ctx context.Context, ent *entity.Role) error
	UpdateMany(ctx context.Context, filter []string, ent map[string]interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter []string) (int64, error)
}

type Transaction interface {
	Count(ctx context.Context, filter []string) (int64, error)
	FindOne(ctx context.Context, filter []string) (*entity.Transaction, error)
	FindMany(ctx context.Context, filter []string) ([]entity.Transaction, error)
	InsertOne(ctx context.Context, ent *entity.Transaction) error
	UpdateMany(ctx context.Context, filter []string, ent map[string]interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter []string) (int64, error)
}

type User interface {
	Count(ctx context.Context, filter []string) (int64, error)
	FindOne(ctx context.Context, filter []string) (*entity.User, error)
	FindMany(ctx context.Context, filter []string) ([]entity.User, error)
	InsertOne(ctx context.Context, ent *entity.User) error
	UpdateMany(ctx context.Context, filter []string, ent map[string]interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter []string) (int64, error)
}
