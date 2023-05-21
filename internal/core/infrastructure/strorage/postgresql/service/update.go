package service

import (
	"context"
	"vm-backend/internal/core/infrastructure/strorage/postgresql"
	"vm-backend/pkg/db"
)

func (r *RepositoryImpl[T]) Update(ctx context.Context, query *db.Query, data map[string]interface{}) (int64, error) {
	var entity T
	tx := postgresql.MakeQuery(r.db.WithContext(ctx).Begin(), query)
	result := tx.Model(&entity).Updates(data)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return result.RowsAffected, nil
}
