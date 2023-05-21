package store

import "time"

type Branch struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	IsEnable  bool      `json:"is_enable"`
}

func (e Branch) TableName() string {
	return "store_branches"
}
