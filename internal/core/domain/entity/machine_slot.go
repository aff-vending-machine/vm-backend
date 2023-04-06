package entity

import (
	"errors"
	"time"
)

type MachineSlot struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Code      string    `json:"code"`
	Stock     int       `json:"stock"`
	Capacity  int       `json:"capacity"`
	MachineID uint      `json:"machine_id"`                                                  // has many
	ProductID uint      `json:"-"`                                                           // belong to
	Product   *Product  `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID"` // belong to
	IsEnable  bool      `json:"is_enable"`
}

func (e MachineSlot) TableName() string {
	return "machine_slots"
}

func (e MachineSlot) Validate() error {
	if e.Code == "" {
		return errors.New("code is required")
	}
	if e.Stock < 0 {
		return errors.New("stock should be greater than or equal to zero")
	}
	if e.Capacity < 0 {
		return errors.New("capacity should be greater than or equal to zero")
	}
	if e.MachineID == 0 {
		return errors.New("machine_id is required")
	}

	return nil
}
