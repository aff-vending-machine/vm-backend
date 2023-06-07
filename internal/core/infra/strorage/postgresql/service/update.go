package service

import (
	"context"
	"vm-backend/internal/core/infra/strorage/postgresql"
	"vm-backend/pkg/helpers/db"
)

func (r *RepositoryImpl[T]) Update(ctx context.Context, query *db.Query, data map[string]interface{}) (int64, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var entity T
	tx := postgresql.MakeQuery(r.db.WithContext(ctx).Begin(), query)
	result := tx.Model(&entity).Updates(data)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}
