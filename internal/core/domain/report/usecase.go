package report

import (
	"context"
	"vm-backend/internal/layer/usecase/report/request"
	"vm-backend/internal/layer/usecase/report/response"
)

type Usecase interface {
	Summary(context.Context, *request.Summary) ([]response.Machine, error)
	Stocks(context.Context, *request.Report) ([]response.Stock, error)
	Transactions(context.Context, *request.Report) ([]response.Transaction, error)
}
