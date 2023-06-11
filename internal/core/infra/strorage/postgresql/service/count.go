package service

import (
	"context"
	"vm-backend/internal/core/infra/strorage/postgresql"
	"vm-backend/pkg/helpers/db"
)

func (r *RepositoryImpl[T]) Count(ctx context.Context, query *db.Query) (int64, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var count int64
	var entity T
	tx := postgresql.MakeQuery(r.db.WithContext(ctx), query)
	result := tx.Model(&entity).Count(&count)
	if err := result.Error; err != nil {
		return 0, err
	}
	
	return count, nil
}
