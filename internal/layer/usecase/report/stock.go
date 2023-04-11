package report

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

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

	transactions, err := uc.transactionRepo.FindMany(ctx, req.ToPaymentFilter())
	if err != nil {
		return nil, errors.Wrap(err, "unable to find transaction")
	}
	if len(transactions) == 0 {
		return []response.Stock{}, nil
	}

	carts := make([]string, len(transactions))
	for i, transaction := range transactions {
		carts[i] = transaction.RawCart
	}

	quantityByCodename := make(map[string]int)
	for _, cartJSON := range carts {
		var cart []response.CartItem
		err := json.Unmarshal([]byte(cartJSON), &cart)
		if err != nil {
			log.Error().Err(err).Msg("unable to unmarshal cart")
			continue
		}

		for _, item := range cart {
			codename := fmt.Sprintf("%s|||%s|||%0.02f", item.Code, item.Name, item.Price)
			quantityByCodename[codename] += item.Quantity
		}
	}
	stocks := make([]response.Stock, 0)

	for codename, quantity := range quantityByCodename {
		cnp := strings.Split(codename, "|||")
		code := cnp[0]
		name := cnp[1]
		price, _ := strconv.ParseFloat(cnp[2], 64)
		stock := response.Stock{
			Code:       code,
			Name:       name,
			Sold:       quantity,
			Price:      price,
			TotalPrice: float64(quantity) * price,
		}

		stocks = append(stocks, stock)
	}

	sort.Sort(response.SortStockByCode(stocks))

	return stocks, nil
}
