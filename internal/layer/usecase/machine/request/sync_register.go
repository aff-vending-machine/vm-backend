package request

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/model"
)

type SyncRegister struct {
	Data model.Machine
}

func (r *SyncRegister) ToFilter() []string {
	return []string{
		fmt.Sprintf("serial_number:=:%s", r.Data.SerialNumber),
	}
}

func (r *SyncRegister) ToEntity() *entity.Machine {
	t := time.Now()
	return &entity.Machine{
		Name:           r.Data.Name,
		SerialNumber:   r.Data.SerialNumber,
		Location:       r.Data.Location,
		Type:           "<auto register>",
		Vendor:         r.Data.Vendor,
		LastActiveTime: &t,
		Status:         "active",
	}
}

func (r *SyncRegister) ToJsonUpdate(count int) map[string]interface{} {
	return map[string]interface{}{
		"last_active_time": time.Now(),
		"status":           "active",
		"count":            count + 1,
	}
}
