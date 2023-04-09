package sync

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/sync/request"
	"github.com/aff-vending-machine/vm-backend/pkg/errs"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) GetTransaction(ctx context.Context, req *request.Sync) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	machine, err := uc.machineRepo.FindOne(ctx, req.ToMachineFilter())
	if err != nil {
		return errors.Wrapf(err, "unable to find machine %d", req.MachineID)
	}

	transactions, err := uc.rpcAPI.GetTransaction(ctx, machine.SerialNumber)
	if err != nil {
		return errors.Wrapf(err, "unable to sync real machine %s", machine.SerialNumber)
	}

	ids := make([]uint, len(transactions))
	for i, transaction := range transactions {
		ids[i] = transaction.ID
		_, err := uc.transactionRepo.FindOne(ctx, []string{fmt.Sprintf("merchant_order_id||=||%s", transaction.MerchantOrderID)})
		if errs.Is(err, "not found") {
			err = uc.transactionRepo.InsertOne(ctx, transaction.ToEntity(machine.ID, machine.Name))
		}
		if err != nil {
			log.Error().Err(err).Msg("unable to find or create transaction")
		}
	}

	uc.rpcAPI.ClearTransaction(ctx, machine.SerialNumber, ids)

	return nil
}
