package usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/response"
)

type Report interface {
	Stock(context.Context, *request.Report) ([]response.Stock, error)
	Transaction(context.Context, *request.Report) ([]response.Transaction, error)
}
