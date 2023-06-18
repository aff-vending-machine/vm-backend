package service

import (
	"context"
	"vm-backend/internal/core/infra/strorage/postgresql"
	"vm-backend/pkg/helpers/db"

	"github.com/rs/zerolog/log"
)

func (r *RepositoryImpl[T]) FindOne(ctx context.Context, query *db.Query) (*T, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var entity T
	tx := r.db.WithContext(ctx)
	tx = postgresql.MakeQuery(tx, query)
	tx = tx.First(&entity)
	if err := tx.Error; err != nil {
		log.Error().Err(err).Msg("unable to find one")
		return nil, err
	}

	return &entity, nil
}
