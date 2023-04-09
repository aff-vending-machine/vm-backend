package transaction

import (
	"context"
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Done(ctx context.Context, req *request.Done) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()

	origin, err := uc.transactionRepo.FindOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "unable to find transaction")
	}

	now := time.Now()
	data := map[string]interface{}{
		"cancelled_by":      nil,
		"cancelled_at":      nil,
		"refund_at":         nil,
		"refund_price":      0,
		"received_item_at":  now,
		"received_quantity": origin.OrderQuantity,
		"paid_price":        origin.PaidPrice,
		"note":              fmt.Sprintf("confirm order by %s", req.Caller),
	}

	if origin.ConfirmedPaidBy == nil {
		data["confirmed_paid_by"] = fmt.Sprintf("user (%s)", req.Caller)
		data["confirmed_paid_at"] = now
	}

	_, err = uc.transactionRepo.UpdateMany(ctx, filter, data)
	if err != nil {
		return errors.Wrap(err, "unable to update transaction")
	}

	return nil
}
