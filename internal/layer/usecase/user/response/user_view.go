package response

import (
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
)

type UserView struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Role      string     `json:"role"`
	LastLogin *time.Time `json:"last_login"`
}

func UserEntityToView(e *entity.User) *UserView {
	return &UserView{
		ID:        e.ID,
		Username:  e.Username,
		Role:      e.Role.Name,
		LastLogin: e.LastLogin,
	}
}

func UserEntityToList(users []entity.User) []UserView {
	items := make([]UserView, len(users))
	for i, user := range users {
		items[i] = *UserEntityToView(&user)
	}

	return items
}
