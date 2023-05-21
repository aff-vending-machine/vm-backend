package response

import (
	"time"
)

type User struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Username  string     `json:"username"`
	Role      string     `json:"role"`
	CreatedBy string     `json:"created_by"`
	LastLogin *time.Time `json:"last_login"`
}
