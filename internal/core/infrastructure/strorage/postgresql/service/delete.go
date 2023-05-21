package service

import (
	"context"
	"vm-backend/internal/core/infrastructure/strorage/postgresql"
	"vm-backend/pkg/db"
)

func (r *RepositoryImpl[T]) Delete(ctx context.Context, query *db.Query) (int64, error) {
	var entity T
	tx := postgresql.MakeQuery(r.db.WithContext(ctx).Begin(), query)
	result := tx.Delete(&entity)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return result.RowsAffected, nil
}
