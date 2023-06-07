package payment_transaction

import (
	"vm-backend/internal/core/infra/network/fiber/http"
	"vm-backend/internal/layer/usecase/payment_transaction/request"

	"github.com/gofiber/fiber/v2"
)

func (r *transportImpl) Cancel(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req, err := makeCancelRequest(c)
	if err != nil {
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.Cancel(ctx, req)
	if err != nil {
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}

func makeCancelRequest(c *fiber.Ctx) (*request.Cancel, error) {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	var req request.Cancel
	req.ID = uint(id)
	req.Caller = getUser(c)

	return &req, nil
}
