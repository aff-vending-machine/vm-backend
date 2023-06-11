package account

import (
	"time"
)

type Token struct {
	ID       uint          `json:"sub"`
	Name     string        `json:"name"`
	RoleID   uint          `json:"role_id"`
	Role     string        `json:"role"`
	BranchID uint          `json:"branch_id"`
	Branch   string        `json:"branch"`
	Type     string        `json:"type"`
	Alive    time.Duration `json:"-,omitempty"`
}

func NewAccessToken(user User) Token {
	branchID := uint(0)
	branchName := "all"

	if user.BranchID != nil {
		branchID = *user.BranchID
		branchName = user.Branch.Name
	}

	return Token{
		ID:       user.ID,
		Name:     user.Username,
		RoleID:   user.RoleID,
		Role:     user.Role.Name,
		BranchID: branchID,
		Branch:   branchName,
		Type:     "ACCESS_TOKEN",
		Alive:    24 * time.Hour,
	}
}

func NewRefreshToken(user User) Token {
	branchID := uint(0)
	branchName := "all"

	if user.BranchID != nil {
		branchID = *user.BranchID
		branchName = user.Branch.Name
	}

	return Token{
		ID:       user.ID,
		Name:     user.Username,
		RoleID:   user.RoleID,
		Role:     user.Role.Name,
		BranchID: branchID,
		Branch:   branchName,
		Type:     "REFRESH_TOKEN",
		Alive:    30 * 24 * time.Hour,
	}
}
