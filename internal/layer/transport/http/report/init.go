package report

import (
	"fmt"
	"time"

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

func generateFilename(req *request.Report, name string) string {
	to := time.Now()
	from := time.Date(2023, time.January, 1, 0, 0, 0, 0, to.Location())
	if req.From != nil {
		from = *req.From
	}
	if req.To != nil {
		to = *req.To
	}

	filename := fmt.Sprintf("%s-%s-%s.csv", name, from.Format("20060102"), to.Format("20060102"))
	return filename
}
