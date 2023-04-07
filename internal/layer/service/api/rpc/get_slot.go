package rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/sync"
	"github.com/aff-vending-machine/vm-backend/pkg/trace"
)

type GetSlotResponse struct {
	Code    int         `json:"code"`
	Data    []sync.Slot `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}

func (r *rpcImpl) GetSlot(ctx context.Context, target string) ([]sync.Slot, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	bres, err := r.EmitRPC(ctx, target, "slot.get", nil)
	if err != nil {
		return nil, err
	}

	var res GetSlotResponse
	err = json.Unmarshal(bres, &res)
	if err != nil {
		return nil, err
	}

	if res.Code != 200 {
		return nil, fmt.Errorf(res.Error)
	}

	return res.Data, nil
}
