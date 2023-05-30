package sync

import (
	"context"
	"time"

	"vm-backend/internal/layer/usecase/sync/request"
	"vm-backend/pkg/db"
	"vm-backend/pkg/errs"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) PullTransactions(ctx context.Context, req *request.Sync) error {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return errors.Wrap(err, "unable to validate request")
	}

	query := req.ToMachineQuery()
	machine, err := uc.machineRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find machine")
		return errors.Wrap(err, "unable to find machine")
	}

	channels, err := uc.channelRepo.FindMany(ctx, db.NewQuery())
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find many channels")
		return errors.Wrap(err, "unable to find many channels")
	}

	channelGroup := make(map[string]uint, len(channels))
	for _, channel := range channels {
		channelGroup[channel.Channel] = channel.ID
	}

	transactions, err := uc.syncAPI.GetTransactions(ctx, machine.SerialNumber)
	if err != nil {
		log.Error().Err(err).Str("target", machine.SerialNumber).Msg("unable to get transactions")
		return errors.Wrap(err, "unable to get transactions")
	}

	ids := make([]uint, len(transactions))
	for i, transaction := range transactions {
		query := db.NewQuery().AddWhere("merchant_order_id", transaction.MerchantOrderID)
		transInDB, err := uc.transactionRepo.FindOne(ctx, query)
		if errs.Is(err, errs.ErrNotFound) {
			channelID := channelGroup[transaction.PaymentChannel]
			_, err = uc.transactionRepo.Create(ctx, transaction.ToDomain(machine.ID, machine.Name, machine.BranchID, channelID))
		}
		if err != nil {
			log.Error().Err(err).Interface("query", query).Msg("unable to find or create transaction")
			continue
		}

		if transInDB != nil && transInDB.OrderStatus != transaction.OrderStatus {
			// updated from vending machine
			if transInDB.OrderStatus != "DONE" && transInDB.OrderStatus != "CANCELLED" {
				update := transaction.ToUpdate()
				_, err := uc.transactionRepo.Update(ctx, query, update)
				if err != nil {
					log.Error().Err(err).Interface("query", query).Interface("update", update).Msg("unable to update transaction")
					continue
				}
			}
		}

		ids[i] = transaction.ID
	}

	if len(ids) > 5 {
		clearedIDs := ids[:len(ids)-5]
		err := uc.syncAPI.ClearTransactions(ctx, machine.SerialNumber, clearedIDs)
		if err != nil {
			log.Error().Err(err).Str("target", machine.SerialNumber).Uints("ids", clearedIDs).Msg("unable to clear transactions")
		}
	}

	query = req.ToMachineQuery()
	update := map[string]interface{}{"sync_transaction_time": time.Now()}
	_, err = uc.machineRepo.Update(ctx, query, update)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Interface("update", update).Msg("unable to update machine")
		return errors.Wrap(err, "unable to update machine")
	}

	return nil
}
