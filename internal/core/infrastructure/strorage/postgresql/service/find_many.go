package service

import (
	"context"
	"vm-backend/internal/core/infrastructure/strorage/postgresql"
	"vm-backend/pkg/db"
)

func (r *RepositoryImpl[T]) FindMany(ctx context.Context, query *db.Query) ([]T, error) {
	var entities []T
	tx := postgresql.MakeQuery(r.db.WithContext(ctx), query)
	result := tx.Find(&entities)
	if err := result.Error; err != nil {
		return nil, err
	}
	return entities, nil
}
