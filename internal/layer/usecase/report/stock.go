package report

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Stock(ctx context.Context, req *request.Report) ([]response.Stock, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	slots, err := uc.machineSlotRepo.FindMany(ctx, req.ToSlotFilter())
	if err != nil {
		return nil, errors.Wrap(err, "unable to find slot")
	}

	transactions, err := uc.transactionRepo.FindMany(ctx, req.ToPaymentFilter())
	if err != nil {
		return nil, errors.Wrap(err, "unable to find transaction")
	}

	carts := make([]string, len(transactions))
	for i, transaction := range transactions {
		carts[i] = transaction.RawCart
	}

	var wg sync.WaitGroup

	quantityByCode := make(map[string]int)
	priceByCode := make(map[string]float64)
	for _, cartJSON := range carts {
		wg.Add(1)

		go func(cartJSON string) {
			defer wg.Done()

			var cart []response.CartItem
			err := json.Unmarshal([]byte(cartJSON), &cart)
			if err != nil {
				log.Error().Err(err).Msg("unable to unmarshal cart")
				return
			}

			for _, item := range cart {
				quantityByCode[item.Code] += item.Quantity
				priceByCode[item.Code] += item.Price * float64(item.Quantity)
			}
		}(cartJSON)
	}
	wg.Wait()

	to := time.Now()
	from := time.Date(2023, time.January, 1, 0, 0, 0, 0, to.Location())
	if req.From != nil {
		from = *req.From
	}
	if req.To != nil {
		to = *req.To
	}

	stocks := make([]response.Stock, len(slots))
	for i, slot := range slots {
		stocks[i] = response.Stock{
			Code:     slot.Code,
			Stock:    slot.Stock,
			Capacity: slot.Capacity,
			Sold:     0,
			From:     from,
			To:       to,
			Price:    0,
		}
	}

	for code, quantity := range quantityByCode {
		for i := range stocks {
			if stocks[i].Code == code {
				stocks[i].Sold = quantity
				stocks[i].Price = priceByCode[code] * float64(quantity)
			}
		}
	}

	if req.Available != nil && *req.Available {
		result := make([]response.Stock, 0)
		for _, stock := range stocks {
			if stock.Sold != 0 {
				result = append(result, stock)
			}
		}
		return result, nil
	}

	return stocks, nil
}
