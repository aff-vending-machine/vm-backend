package report

import (
	"vm-backend/internal/core/domain/report"
	"vm-backend/internal/layer/usecase/report/request"

	"github.com/gofiber/fiber/v2"
)

type transportImpl struct {
	usecase report.Usecase
}

func NewTransport(uc report.Usecase) report.Transport {
	return &transportImpl{uc}
}

func makeReportRequest(c *fiber.Ctx) (*request.Report, error) {
	machineID, err := c.ParamsInt("machine_id", 0)
	if err != nil {
		return nil, err
	}

	var req request.Report
	if err := c.QueryParser(&req); err != nil {
		return nil, err
	}
	req.MachineID = uint(machineID)

	return &req, nil
}
