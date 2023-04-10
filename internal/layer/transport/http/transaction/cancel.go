package transaction

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/transaction/request"
	"github.com/aff-vending-machine/vm-backend/pkg/trace"
	"github.com/gofiber/fiber/v2"
)

func (r *httpImpl) Cancel(c *fiber.Ctx) error {
	ctx, span := trace.Start(c.Context())
	defer span.End()

	req, err := makeCancelRequest(c)
	if err != nil {
		trace.RecordError(span, err)
		return http.BadRequest(c, err)
	}

	// usecase execution
	err = r.usecase.Cancel(ctx, req)
	if err != nil {
		trace.RecordError(span, err)
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
