package sync

import (
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type Machine struct {
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
	Branch       string `json:"branch"`
	Location     string `json:"location"`
	Type         string `json:"type"`
	Vendor       string `json:"vendor"`
	Center       string `json:"center"`
	Status       string `json:"status"`
}

func ToMachine(e *entity.Machine) *Machine {
	return &Machine{
		Name:         e.Name,
		SerialNumber: e.SerialNumber,
		Branch:       e.Branch,
		Location:     e.Location,
		Type:         e.Type,
		Vendor:       e.Vendor,
		Status:       e.Status,
	}
}

func (m *Machine) ToJsonUpdate() map[string]interface{} {
	return map[string]interface{}{
		"name":              m.Name,
		"serial_number":     m.SerialNumber,
		"branch":            m.Branch,
		"location":          m.Location,
		"type":              m.Type,
		"vendor":            m.Vendor,
		"center":            m.Center,
		"status":            m.Status,
		"sync_machine_time": time.Now(),
	}
}
