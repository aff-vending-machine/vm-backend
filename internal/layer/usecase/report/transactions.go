package report

import (
	"context"
	"encoding/json"
	"time"

	"vm-backend/internal/layer/usecase/report/request"
	"vm-backend/internal/layer/usecase/report/response"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Transactions(ctx context.Context, req *request.Report) ([]response.Transaction, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToTransactionQuery()
	trans, err := uc.transactionRepo.FindMany(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find many transactions")
		return nil, errors.Wrap(err, "unable to find many transactions")
	}

	res := make([]response.Transaction, len(trans))
	for i, t := range trans {
		res[i] = response.Transaction{
			MerchantOrderID:     t.MerchantOrderID,
			MachineID:           t.MachineID,
			MachineName:         t.Machine.Name,
			MachineSerialNumber: t.Machine.SerialNumber,
			Location:            t.Machine.Location,
			Cart:                []response.Item{},
			PaymentChannel:      t.Channel.Channel,
			ConfirmedPaidBy:     "",
			OrderedAt:           t.OrderedAt,
			PaymentRequestedAt:  time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
			ConfirmedPaidAt:     time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
			ReceivedItemAt:      time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
			OrderQuantity:       t.OrderQuantity,
			ReceivedQuantity:    t.ReceivedQuantity,
			OrderPrice:          t.OrderPrice,
			PaidPrice:           t.PaidPrice,
			Reference1:          "",
			Reference2:          "",
			Reference3:          "",
			Note:                t.Note,
		}

		if t.ConfirmedPaidBy != nil {
			res[i].ConfirmedPaidBy = *t.ConfirmedPaidBy
		}

		if t.PaymentRequestedAt != nil {
			res[i].PaymentRequestedAt = *t.PaymentRequestedAt
		}

		if t.ConfirmedPaidAt != nil {
			res[i].ConfirmedPaidAt = *t.ConfirmedPaidAt
		}

		if t.ReceivedItemAt != nil {
			res[i].ReceivedItemAt = *t.ReceivedItemAt
		}

		if t.Reference1 != nil {
			res[i].Reference1 = *t.Reference1
		}

		if t.Reference2 != nil {
			res[i].Reference2 = *t.Reference2
		}

		if t.Reference3 != nil {
			res[i].Reference3 = *t.Reference3
		}

		err := json.Unmarshal([]byte(t.RawCart), &res[i].Cart)
		if err != nil {
			log.Warn().Err(err).Interface("cart", t.RawCart).Msg("unable to unmarshal cart")
		}
	}

	return res, nil
}
