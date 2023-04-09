package report_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (r *httpImpl) GetTransaction(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeReportRequest(c)
	if err != nil {
		log.Error().Err(err).Msg("unable to make report request")
		return http.BadRequest(c, err)
	}

	// usecase execution
	res, err := r.usecase.Transaction(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("unable to get transaction report")
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}
