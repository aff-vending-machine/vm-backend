package sync_rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"vm-backend/internal/core/domain/sync"
	"vm-backend/internal/core/domain/sync/models"
)

func (r *rpcImpl) GetTransactions(ctx context.Context, target string) ([]models.Transaction, error) {
	bres, err := r.EmitRPC(ctx, target, "transaction.get", nil)
	if err != nil {
		return nil, err
	}

	var res sync.Response[[]models.Transaction]
	err = json.Unmarshal(bres, &res)
	if err != nil {
		return nil, err
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf(*res.Error)
	}

	return *res.Data, nil
}
