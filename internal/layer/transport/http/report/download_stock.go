package report

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (r *httpImpl) DownloadStock(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeReportRequest(c)
	if err != nil {
		log.Error().Err(err).Msg("unable to make report request")
		return http.BadRequest(c, err)
	}

	// usecase execution
	res, err := r.usecase.Stock(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("unable to get stock report")
		return http.UsecaseError(c, err)
	}

	csv, err := r.usecase.StockCSV(ctx, req, res)
	if err != nil {
		log.Error().Err(err).Msg("unable to create stock csv report")
		return http.UsecaseError(c, err)
	}

	name := generateFilename(req, "stock")

	return http.SendFile(c, name, csv)
}
