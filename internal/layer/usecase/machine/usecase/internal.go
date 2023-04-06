package machine_usecase

import (
	"context"
	"time"
)

func (uc *usecaseImpl) updateMachineStatus(ctx context.Context, filter []string, status string) {
	data := map[string]interface{}{
		"status": status,
	}

	switch status {
	case "active":
		data["last_active_time"] = time.Now()
	case "maintenance":
		data["last_maintenance_time"] = time.Now()
	}

	uc.machineRepo.UpdateMany(ctx, filter, data)
}
