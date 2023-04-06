package entity

import (
	"errors"
	"time"
)

type Machine struct {
	ID                  uint          `json:"id" gorm:"primarykey"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
	Name                string        `json:"name"`
	SerialNumber        string        `json:"serial_number" gorm:"uniqueIndex"`
	Branch              string        `json:"branch"`
	Location            string        `json:"location"`
	Type                string        `json:"type"`
	Vendor              string        `json:"vendor"`
	Status              string        `json:"status"`
	LastActiveTime      *time.Time    `json:"last_active_time"`
	LastMaintenanceTime *time.Time    `json:"last_maintenance_time"`
	Slots               []MachineSlot `json:"slots" gorm:"foreignKey:MachineID"`
}

func (e Machine) TableName() string {
	return "machines"
}

func (e Machine) Validate() error {
	if e.Name == "" {
		return errors.New("name is required")
	}
	if e.SerialNumber == "" {
		return errors.New("serial_number is required")
	}
	if e.Status == "" {
		return errors.New("status is required")
	}

	return nil
}
