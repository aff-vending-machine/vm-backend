package payment_transaction

import (
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/core/infra/network/fiber/http"
	"vm-backend/internal/layer/usecase/payment_transaction/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) Done(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeDoneRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.Done(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}

func makeDoneRequest(c *fiber.Ctx) (*request.Done, error) {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	var req request.Done
	req.ID = uint(id)
	req.Caller = account.GetAccess(c)
	req.BranchID = account.GetBranchID(c, req.BranchID)

	return &req, nil
}
