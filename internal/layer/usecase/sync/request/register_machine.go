package request

import (
	"time"
	"vm-backend/internal/core/domain/sync/models"
	"vm-backend/pkg/db"
)

type RegisterMachine struct {
	Data models.Machine `json:"data"`
}

func (r *RegisterMachine) ToQuery() *db.Query {
	return db.NewQuery().AddWhere("serial_number", r.Data.SerialNumber)
}

func (r *RegisterMachine) ToUpdate(count int) map[string]interface{} {
	return map[string]interface{}{
		"sync_time": time.Now(),
		"status":    "enable",
		"count":     count + 1,
	}
}
