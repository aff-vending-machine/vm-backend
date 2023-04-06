package rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/model"
	"github.com/aff-vending-machine/vm-backend/pkg/trace"
)

type SlotSetRequest struct {
	Data []model.Slot `json:"data"`
}

type SlotSetResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (r *rpcImpl) SlotSet(ctx context.Context, target string, data []model.Slot) error {
	_, span := trace.Start(ctx)
	defer span.End()

	req := SlotSetRequest{Data: data}
	breq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	bres, err := r.Emit(ctx, target, "slot.set", breq)
	if err != nil {
		return err
	}

	var res SlotSetResponse
	err = json.Unmarshal(bres, &res)
	if err != nil {
		return err
	}

	if res.Code != 200 {
		return fmt.Errorf(res.Error)
	}

	return nil
}
