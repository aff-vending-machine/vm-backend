package entity

import (
	"time"
)

type Permission struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RoleID    uint      `json:"role_id"`
	Scope     string    `json:"scope"`
	Level     int       `json:"level"`
}
