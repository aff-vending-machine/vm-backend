package sync_rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"vm-backend/internal/core/domain/sync"
	"vm-backend/internal/core/domain/sync/models"
)

func (r *rpcImpl) SetSlots(ctx context.Context, target string, data []models.Slot) error {
	req := sync.Request[struct{}, []models.Slot]{Data: &data}

	bres, err := r.EmitRPC(ctx, target, "slot.set", req.ToBytes())
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
