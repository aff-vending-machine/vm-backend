package service

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (r *RepositoryImpl[T]) Create(ctx context.Context, entity *T) (uint, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	tx := r.db.WithContext(ctx).Begin()
	tx = tx.Create(entity)
	if err := tx.Error; err != nil {
		log.Error().Err(err).Msg("unable to create entity")
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit().Error; err != nil {
		log.Error().Err(err).Msg("unable to commit transaction")
		return 0, err
	}
	
	return uint(tx.RowsAffected), nil
}
