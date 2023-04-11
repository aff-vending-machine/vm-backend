package rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/sync"
	"github.com/aff-vending-machine/vm-backend/pkg/trace"
)

type SetMachineRequest struct {
	Data sync.Machine `json:"data"`
}

type SetMachineResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (r *rpcImpl) SetMachine(ctx context.Context, target string, data *sync.Machine) error {
	_, span := trace.Start(ctx)
	defer span.End()

	req := SetMachineRequest{Data: *data}
	breq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	bres, err := r.EmitRPC(ctx, target, "machine.set", breq)
	if err != nil {
		return err
	}

	var res SetMachineResponse
	err = json.Unmarshal(bres, &res)
	if err != nil {
		return err
	}

	if res.Code != 200 {
		return fmt.Errorf(res.Error)
	}

	return nil
}
