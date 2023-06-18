package migration

import (
	"context"
	"vm-backend/internal/boot/modules"
	"vm-backend/pkg/helpers/db"

	"github.com/rs/zerolog/log"
)

func UpdateBranchInTransactions(repo modules.RepositoryService) {
	log.Debug().Msg("updating branch in transactions")
	ctx := context.Background()
	machines, err := repo.Machine.FindMany(ctx, db.NewQuery())
	if err != nil {
		log.Error().Err(err).Msg("unable to find transactions")
		return
	}

	branches, err := repo.StoreBranch.FindMany(ctx, db.NewQuery())
	if err != nil {
		log.Error().Err(err).Msg("unable to find branches")
		return
	}

	locationMap := make(map[string]uint)
	for _, branch := range branches {
		locationMap[branch.Location] = branch.ID
	}

	for _, machine := range machines {
		branchID := locationMap[machine.Location]

		total, err := repo.PaymentTransaction.Update(ctx, db.NewQuery().Where("machine_id = ?", machine.ID).Where("branch_id = ?", 1), map[string]interface{}{"branch_id": branchID})
		if err != nil {
			log.Error().Err(err).Uint("id", machine.ID).Uint("branch_id", branchID).Msg("unable to update transaction")
			return
		}
		log.Info().Uint("id", machine.ID).Uint("branch_id", branchID).Int64("total", total).Msg("transaction updated")
	}
}
