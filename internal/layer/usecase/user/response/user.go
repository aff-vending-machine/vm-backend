package response

import (
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
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

func ToUser(e *entity.User) *User {
	return &User{
		ID:        e.ID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		Username:  e.Username,
		Role:      e.Role.Name,
		CreatedBy: e.CreatedBy,
		LastLogin: e.LastLogin,
	}
}

func ToUserList(users []entity.User) []User {
	items := make([]User, len(users))
	for i, user := range users {
		items[i] = *ToUser(&user)
	}

	return items
}
