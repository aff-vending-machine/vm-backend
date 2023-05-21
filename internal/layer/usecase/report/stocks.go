package report

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"vm-backend/internal/layer/usecase/report/request"
	"vm-backend/internal/layer/usecase/report/response"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

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
			codename := fmt.Sprintf("%s:%s:%0.02f", item.Code, item.Name, item.Price)
			quantityByCodename[codename] += item.Quantity
		}
	}
	stocks := make([]response.Stock, 0)

	for codename, quantity := range quantityByCodename {
		cnp := strings.Split(codename, ":")
		code := cnp[0]
		name := cnp[1]
		price, _ := strconv.ParseFloat(cnp[2], 64)
		stock := response.Stock{
			Code:       code,
			Name:       name,
			Sold:       quantity,
			SalePrice:  price,
			TotalPrice: float64(quantity) * price,
		}

		stocks = append(stocks, stock)
	}

	sort.Sort(response.SortStockByCode(stocks))

	return stocks, nil
}
