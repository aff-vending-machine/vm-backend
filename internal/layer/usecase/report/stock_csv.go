package report

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/response"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) StockCSV(ctx context.Context, req *request.Report, records []response.Stock) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)

	header := []string{"Code", "Product", "Sold", "Price", "Total Price"}
	err := writer.Write(header)
	if err != nil {
		log.Error().Err(err).Strs("headers", header).Msg("unable to write headers")
		return nil, err
	}

	for _, record := range records {
		data := []string{
			record.Code,
			record.Name,
			fmt.Sprintf("%d", record.Sold),
			fmt.Sprintf("%0.02f", record.Price),
			fmt.Sprintf("%0.02f", record.TotalPrice),
		}
		err = writer.Write(data)
		if err != nil {
			log.Error().Err(err).Strs("data", data).Msg("unable to write data")
		}
	}

	writer.Flush()

	return buf, nil
}
