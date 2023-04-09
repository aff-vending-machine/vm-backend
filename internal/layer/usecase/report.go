package usecase

import (
	"bytes"
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/response"
)

type Report interface {
	Stock(context.Context, *request.Report) ([]response.Stock, error)
	Payment(context.Context, *request.Report) ([]response.Payment, error)
	StockCSV(context.Context, *request.Report, []response.Stock) (*bytes.Buffer, error)
	PaymentCSV(context.Context, *request.Report, []response.Payment) (*bytes.Buffer, error)
}
