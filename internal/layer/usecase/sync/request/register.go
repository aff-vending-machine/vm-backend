package request

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/sync"
)

type RegisterMachine struct {
	Data sync.Machine `json:"data"`
}

func (r *RegisterMachine) ToFilter() []string {
	return []string{
		fmt.Sprintf("serial_number:=:%s", r.Data.SerialNumber),
	}
}

func (r *RegisterMachine) ToEntity() *entity.Machine {
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

func (r *RegisterMachine) ToJsonUpdate(count int) map[string]interface{} {
	return map[string]interface{}{
		"last_active_time": time.Now(),
		"status":           "active",
		"count":            count + 1,
	}
}
