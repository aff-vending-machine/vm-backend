package service

import (
	"context"
	"vm-backend/internal/core/infrastructure/strorage/postgresql"
	"vm-backend/pkg/db"
)

func (r *RepositoryImpl[T]) FindOne(ctx context.Context, query *db.Query) (*T, error) {
	var entity T
	tx := postgresql.MakeQuery(r.db.WithContext(ctx), query)
	result := tx.First(&entity)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &entity, nil
}
