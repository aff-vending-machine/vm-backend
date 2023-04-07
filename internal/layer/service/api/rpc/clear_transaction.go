package rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/pkg/trace"
)

type ClearTransactionRequest struct {
	Query ClearTransactionQuery `json:"query"`
}

type ClearTransactionQuery struct {
	IDs []uint `json:"ids"`
}

type ClearTransactionResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (r *rpcImpl) ClearTransaction(ctx context.Context, target string, ids []uint) error {
	_, span := trace.Start(ctx)
	defer span.End()

	req := ClearTransactionRequest{Query: ClearTransactionQuery{IDs: ids}}
	breq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	bres, err := r.EmitRPC(ctx, target, "transaction.clear", breq)
	if err != nil {
		return err
	}

	var res ClearTransactionResponse
	err = json.Unmarshal(bres, &res)
	if err != nil {
		return err
	}

	if res.Code != 200 {
		return fmt.Errorf(res.Error)
	}

	return nil
}
