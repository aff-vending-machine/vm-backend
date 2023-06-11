package response

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	RoleID   uint   `json:"role_id"`
	Role     string `json:"role"`
	BranchID uint   `json:"branch_id"`
	Branch   string `json:"branch"`
	Type     string `json:"type"`
}
