package sync_rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"vm-backend/internal/core/domain/sync"
	"vm-backend/internal/core/domain/sync/models"
)

func (r *rpcImpl) GetChannels(ctx context.Context, target string) ([]models.Channel, error) {
	bres, err := r.EmitRPC(ctx, target, "channel.get", nil)
	if err != nil {
		return nil, err
	}

	var res sync.Response[[]models.Channel]
	err = json.Unmarshal(bres, &res)
	if err != nil {
		return nil, err
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf(*res.Error)
	}

	return *res.Data, nil
}
