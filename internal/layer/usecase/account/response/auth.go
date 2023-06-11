package response

type AuthResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

type User struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	RoleID   uint   `json:"role_id"`
	Role     string `json:"role"`
	BranchID uint   `json:"branch_id"`
	Branch   string `json:"branch"`
}
