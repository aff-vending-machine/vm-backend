package models

import (
	"time"
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

func (m *Machine) ToUpdate() map[string]interface{} {
	return map[string]interface{}{
		"name":          m.Name,
		"serial_number": m.SerialNumber,
		"location":      m.Location,
		"type":          m.Type,
		"vendor":        m.Vendor,
		"center":        m.Center,
		"status":        m.Status,
		"sync_time":     time.Now(),
	}
}
