package migration

import (
	"context"
	"vm-backend/internal/boot/modules"
	"vm-backend/pkg/helpers/db"

	"github.com/rs/zerolog/log"
)

func UpdateAlternativeScope(repo modules.RepositoryService) {
	log.Info().Msg("updating alternative scope")
	ctx := context.Background()
	permissions, err := repo.AccountPermission.FindMany(ctx, db.NewQuery())
	if err != nil {
		log.Error().Err(err).Msg("unable to find permissions")
		return
	}

	for _, permission := range permissions {
		if permission.AltScope != "" {
			// permissions has alternative scope, skip
			continue
		}

		alt := ""
		if permission.Scope == "branch" {
			alt = permission.Scope + "es"
		} else {
			alt = permission.Scope + "s"
		}

		total, err := repo.AccountPermission.Update(ctx, db.NewQuery().Where("id = ?", permission.ID), map[string]interface{}{"alt_scope": alt})
		if err != nil {
			log.Error().Err(err).Uint("id", permission.ID).Str("alt_scope", alt).Msg("unable to update permission")
			return
		}
		log.Info().Uint("id", permission.ID).Str("alt_scope", alt).Int64("total", total).Msg("permission updated")
	}
}
