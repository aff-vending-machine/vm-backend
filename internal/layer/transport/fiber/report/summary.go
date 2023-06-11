package report

import (
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/core/infra/network/fiber/http"
	"vm-backend/internal/layer/usecase/report/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) Summary(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeSummaryRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	res, err := r.usecase.Summary(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}

func makeSummaryRequest(c *fiber.Ctx) (*request.Summary, error) {
	var req request.Summary
	if err := c.QueryParser(&req); err != nil {
		return nil, err
	}
	req.BranchID = account.GetBranchID(c, req.BranchID)

	return &req, nil
}
