package machine_http

import (
	"github.com/aff-vending-machine/vm-backend/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/machine/request"
	"github.com/aff-vending-machine/vm-backend/pkg/trace"
	"github.com/gofiber/fiber/v2"
)

func (r *restImpl) SyncGet(c *fiber.Ctx) error {
	ctx, span := trace.Start(c.Context())
	defer span.End()

	req, err := makeSyncRequest(c)
	if err != nil {
		trace.RecordError(span, err)
		return http.BadRequest(c, err)
	}

	// usecase execution
	res, err := r.usecase.SyncGet(ctx, req)
	if err != nil {
		trace.RecordError(span, err)
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}

func makeSyncRequest(c *fiber.Ctx) (*request.Sync, error) {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}

	return &request.Sync{ID: uint(id)}, nil
}
