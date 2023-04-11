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
	Count               int           `json:"count"`
	SyncMachineTime     *time.Time    `json:"sync_machine_time"`
	SyncSlotTime        *time.Time    `json:"sync_slot_time"`
	SyncTransactionTime *time.Time    `json:"sync_transaction_time"`
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
