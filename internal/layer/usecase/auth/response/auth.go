package response

import "github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"

type AuthResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func ToUser(user *entity.User) User {
	return User{
		Username: user.Username,
		Role:     user.Role.Name,
	}
}
