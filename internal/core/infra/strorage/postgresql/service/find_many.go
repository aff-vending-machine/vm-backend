package service

import (
	"context"
	"vm-backend/internal/core/infra/strorage/postgresql"
	"vm-backend/pkg/helpers/db"

	"github.com/rs/zerolog/log"
)

func (r *RepositoryImpl[T]) FindMany(ctx context.Context, query *db.Query) ([]T, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var entities []T
	tx := r.db.WithContext(ctx)
	tx = postgresql.MakeQuery(tx, query)
	tx = tx.Find(&entities)
	if err := tx.Error; err != nil {
		log.Error().Err(err).Msg("unable to find entities")
		return nil, err
	}

	return entities, nil
}
