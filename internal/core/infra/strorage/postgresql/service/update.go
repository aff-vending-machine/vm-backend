package service

import (
	"context"
	"vm-backend/internal/core/infra/strorage/postgresql"
	"vm-backend/pkg/helpers/db"

	"github.com/rs/zerolog/log"
)

func (r *RepositoryImpl[T]) Update(ctx context.Context, query *db.Query, data map[string]interface{}) (int64, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var entity T
	tx := r.db.WithContext(ctx).Begin()
	tx = postgresql.MakeQuery(tx, query)
	tx = tx.Model(&entity).Updates(data)
	if err := tx.Error; err != nil {
		log.Error().Err(err).Msg("unable to update")
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit().Error; err != nil {
		log.Error().Err(err).Msg("unable to commit transaction")
		return 0, err
	}

	return tx.RowsAffected, nil
}
