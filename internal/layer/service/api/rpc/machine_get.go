package rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/model"
	"github.com/aff-vending-machine/vm-backend/pkg/trace"
)

type MachineGetResponse struct {
	Code    int            `json:"code"`
	Data    *model.Machine `json:"data"`
	Message string         `json:"message"`
	Error   string         `json:"error"`
}

func (r *rpcImpl) MachineGet(ctx context.Context, target string) (*model.Machine, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	bres, err := r.Emit(ctx, target, "slot.get", nil)
	if err != nil {
		return nil, err
	}

	var res MachineGetResponse
	err = json.Unmarshal(bres, &res)
	if err != nil {
		return nil, err
	}

	if res.Code != 200 {
		return nil, fmt.Errorf(res.Error)
	}

	return res.Data, nil
}
