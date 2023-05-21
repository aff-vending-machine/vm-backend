package response

import (
	"time"
)

type Slot struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	MachineID uint      `json:"machine_id"`
	ProductID uint      `json:"product_id"`
	Product   *Product  `json:"product,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Code      string    `json:"code"`
	Stock     int       `json:"stock"`
	Capacity  int       `json:"capacity"`
	IsEnable  bool      `json:"is_enable"`
}
