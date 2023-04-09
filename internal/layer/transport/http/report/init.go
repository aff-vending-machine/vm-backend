package report

import (
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/report/request"
	"github.com/gofiber/fiber/v2"
)

type httpImpl struct {
	usecase usecase.Report
}

func New(uc usecase.Report) *httpImpl {
	return &httpImpl{uc}
}

func makeReportRequest(c *fiber.Ctx) (*request.Report, error) {
	var req request.Report
	if err := c.QueryParser(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
