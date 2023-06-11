package service

import (
	"context"
	"vm-backend/internal/core/infra/strorage/postgresql"
	"vm-backend/pkg/helpers/db"
)

func (r *RepositoryImpl[T]) FindOne(ctx context.Context, query *db.Query) (*T, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var entity T
	tx := postgresql.MakeQuery(r.db.WithContext(ctx), query)
	result := tx.First(&entity)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &entity, nil
}
