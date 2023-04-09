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

func (uc *usecaseImpl) PaymentCSV(ctx context.Context, req *request.Report, records []response.Payment) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)

	header := []string{
		"MerchantOrderID",
		"MachineID",
		"MachineName",
		"MachineSerialNumber",
		"Location",
		"PaymentChannel",
		"ConfirmedPaidBy",
		"OrderedAt",
		"PaymentRequestedAt",
		"ConfirmedPaidAt",
		"ReceivedItemAt",
		"OrderQuantity",
		"ReceivedQuantity",
		"OrderPrice",
		"PaidPrice",
		"Reference1",
		"Reference2",
		"Reference3",
		"Note",
	}
	err := writer.Write(header)
	if err != nil {
		log.Error().Err(err).Strs("headers", header).Msg("unable to write headers")
		return nil, err
	}

	for _, record := range records {
		data := createData(record)
		err = writer.Write(data)
		if err != nil {
			log.Error().Err(err).Strs("data", data).Msg("unable to write data")
		}
	}

	writer.Flush()

	return buf, nil
}

func createData(record response.Payment) []string {
	data := []string{
		record.MerchantOrderID,
		fmt.Sprintf("%d", record.MachineID),
		record.MachineName,
		record.MachineSerialNumber,
		record.Location,
		record.PaymentChannel,
		"",
		record.OrderedAt.Format(TIME_LAYOUT),
		"",
		"",
		"",
		fmt.Sprintf("%d", record.OrderQuantity),
		fmt.Sprintf("%d", record.ReceivedQuantity),
		fmt.Sprintf("%0.02f", record.OrderPrice),
		fmt.Sprintf("%0.02f", record.PaidPrice),
		"",
		"",
		"",
		record.Note,
	}

	if record.ConfirmedPaidBy != nil {
		data[5] = *record.ConfirmedPaidBy
	}
	if record.PaymentRequestedAt != nil {
		data[7] = record.PaymentRequestedAt.Format(TIME_LAYOUT)
	}
	if record.ConfirmedPaidAt != nil {
		data[8] = record.ConfirmedPaidAt.Format(TIME_LAYOUT)
	}
	if record.ReceivedItemAt != nil {
		data[9] = record.ReceivedItemAt.Format(TIME_LAYOUT)
	}
	if record.Reference1 != nil {
		data[14] = *record.Reference1
	}
	if record.Reference2 != nil {
		data[15] = *record.Reference2
	}
	if record.Reference3 != nil {
		data[16] = *record.Reference3
	}

	return data
}
