package report

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"vm-backend/internal/layer/usecase/report/request"
	"vm-backend/internal/layer/usecase/report/response"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type StockData struct {
	Code     string
	Name     string
	Price    float64
	Quantity int
	Payments map[string]float64
}

func (uc *usecaseImpl) Stocks(ctx context.Context, req *request.Report) ([]response.Stock, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToTransactionQuery()
	transactions, err := uc.transactionRepo.FindMany(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find transactions")
		return nil, errors.Wrap(err, "unable to find transactions")
	}
	if len(transactions) == 0 {
		log.Info().Interface("query", query).Msg("no transactions found")
		return []response.Stock{}, nil
	}

	channels := make([]string, len(transactions))
	carts := make([]string, len(transactions))
	for i, transaction := range transactions {
		carts[i] = transaction.RawCart
		channels[i] = transaction.Channel.Channel
	}

	dataByCodename := make(map[string]StockData)
	for i, cartJSON := range carts {
		var cart []response.CartItem
		err := json.Unmarshal([]byte(cartJSON), &cart)
		if err != nil {
			log.Error().Err(err).Msg("unable to unmarshal cart")
			continue
		}

		channel := channels[i]

		for _, item := range cart {
			codename := fmt.Sprintf("%s:%s:%0.02f", item.Code, item.Name, item.Price)
			quantity := dataByCodename[codename].Quantity + item.Quantity
			payments := dataByCodename[codename].Payments
			payments[channel] += float64(item.Quantity) * item.Price

			dataByCodename[codename] = StockData{
				Code:     item.Code,
				Name:     item.Name,
				Price:    item.Price,
				Quantity: quantity,
				Payments: payments,
			}
		}
	}

	stocks := make([]response.Stock, 0)

	for _, data := range dataByCodename {
		stock := response.Stock{
			Code:          data.Code,
			Name:          data.Name,
			Sold:          data.Quantity,
			SalePrice:     data.Price,
			TotalPayments: data.Payments,
			TotalPrice:    float64(data.Quantity) * data.Price,
		}

		stocks = append(stocks, stock)
	}

	sort.Sort(response.SortStockByCode(stocks))

	return stocks, nil
}
