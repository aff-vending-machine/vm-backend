package migration

import (
	"context"
	"vm-backend/internal/boot/modules"
	"vm-backend/internal/core/domain/store"
	"vm-backend/pkg/helpers/db"
	"vm-backend/pkg/helpers/errs"

	"github.com/rs/zerolog/log"
)

func CreateBranchFromMachine(repo modules.RepositoryService) {
	ctx := context.Background()
	machines, err := repo.Machine.FindMany(ctx, db.NewQuery())
	if err != nil {
		log.Error().Err(err).Msg("unable to find machines")
		return
	}

	for _, machine := range machines {
		branchLocation := machine.Location

		if machine.BranchID != nil {
			// machine has branch id, skip
			continue
		}

		branch, err := repo.StoreBranch.FindOne(ctx, db.NewQuery().AddWhere("location", branchLocation))
		if errs.Not(err, errs.ErrNotFound) {
			log.Error().Err(err).Str("location", branchLocation).Msg("unable to find branch")
			return
		}

		if branch == nil {
			// no branch found, create one
			branch = &store.Branch{
				Name:     branchLocation,
				Location: branchLocation,
				IsEnable: true,
			}
			_, err = repo.StoreBranch.Create(ctx, branch)
			if err != nil {
				log.Error().Err(err).Str("location", branchLocation).Msg("unable to create branch")
				return
			}
			log.Info().Str("location", branchLocation).Msg("branch created")
		}

		if machine.BranchID != nil && *machine.BranchID == branch.ID {
			// machine is matched with branch and location, skip
			continue
		}

		_, err = repo.Machine.Update(ctx, db.NewQuery().AddWhere("id = ?", machine.ID), map[string]interface{}{"branch_id": branch.ID})
		if err != nil {
			log.Error().Err(err).Uint("id", machine.ID).Uint("branch_id", branch.ID).Msg("unable to update machine")
			return
		}
		log.Info().Uint("id", machine.ID).Uint("branch_id", branch.ID).Msg("machine updated")
	}
}
