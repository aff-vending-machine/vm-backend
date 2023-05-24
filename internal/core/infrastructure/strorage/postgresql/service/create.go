package service

import (
	"context"
)

func (r *RepositoryImpl[T]) Create(ctx context.Context, entity *T) (uint, error) {
	tx := r.db.WithContext(ctx).Begin()
	result := tx.Create(entity)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}
	return uint(result.RowsAffected), nil
}
