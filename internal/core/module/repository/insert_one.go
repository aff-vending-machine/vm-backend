package repository

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/pkg/trace"
)

func (r *Template[T]) InsertOne(ctx context.Context, data *T) error {
	_, span := trace.Start(ctx)
	defer span.End()

	tx := r.DB.Begin()

	result := tx.Create(data)
	if err := result.Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
