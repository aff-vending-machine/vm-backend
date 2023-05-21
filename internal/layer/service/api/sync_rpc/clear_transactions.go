package sync_rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"vm-backend/internal/core/domain/sync"
)

func (r *rpcImpl) ClearTransactions(ctx context.Context, target string, ids []uint) error {
	type Query struct {
		IDs []uint `json:"ids"`
	}

	req := sync.Request[Query, struct{}]{
		Query: &Query{IDs: ids},
	}

	bres, err := r.EmitRPC(ctx, target, "transaction.clear", req.ToBytes())
	if err != nil {
		return err
	}

	var res sync.Response[struct{}]
	err = json.Unmarshal(bres, &res)
	if err != nil {
		return err
	}

	if res.IsSuccess() {
		return fmt.Errorf(*res.Error)
	}

	return nil
}
