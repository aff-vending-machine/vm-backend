package rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/sync"
	"github.com/aff-vending-machine/vm-backend/pkg/trace"
)

type GetTransactionResponse struct {
	Code    int                `json:"code"`
	Data    []sync.Transaction `json:"data"`
	Message string             `json:"message"`
	Error   string             `json:"error"`
}

func (r *rpcImpl) GetTransaction(ctx context.Context, target string) ([]sync.Transaction, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	bres, err := r.EmitRPC(ctx, target, "transaction.get", nil)
	if err != nil {
		return nil, err
	}

	var res GetTransactionResponse
	err = json.Unmarshal(bres, &res)
	if err != nil {
		return nil, err
	}

	if res.Code != 200 {
		return nil, fmt.Errorf(res.Error)
	}

	return res.Data, nil
}
