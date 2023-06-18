package service

import (
	"context"
	"vm-backend/internal/core/infra/strorage/postgresql"
	"vm-backend/pkg/helpers/db"

	"github.com/rs/zerolog/log"
)

func (r *RepositoryImpl[T]) Count(ctx context.Context, query *db.Query) (int64, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var count int64
	var entity T
	tx := r.db.WithContext(ctx)
	tx = postgresql.MakeQuery(tx, query)
	tx = tx.Model(&entity).Count(&count)
	if err := tx.Error; err != nil {
		log.Error().Err(err).Msg("unable to count")
		return 0, err
	}

	return count, nil
}
