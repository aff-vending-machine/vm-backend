package transaction

import (
	"context"
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Cancel(ctx context.Context, req *request.Cancel) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	filter := req.ToFilter()

	data := map[string]interface{}{
		"order_status":      "CANCELLED",
		"cancelled_by":      req.Caller,
		"cancelled_at":      time.Now(),
		"refund_at":         nil,
		"refund_price":      0,
		"received_item_at":  nil,
		"received_quantity": 0,
		"paid_price":        0,
		"note":              fmt.Sprintf("cancel order by %s", req.Caller),
		"confirmed_paid_by": nil,
		"confirmed_paid_at": nil,
	}

	_, err := uc.transactionRepo.UpdateMany(ctx, filter, data)
	if err != nil {
		return errors.Wrap(err, "unable to update transaction")
	}

	return nil
}
